package main

import (
	"html/template"
	"net/http"
	"strings"
)

// Load and parse all HTML templates inside the templates folder
var tmpl = template.Must(template.ParseGlob("templates/*.html"))

// Structure used to send data to HTML templates
type pageData struct {
	Result string // Generated ASCII art
	Text   string // User input text
	Banner string // Selected banner style
}

// Handles GET requests to the home page (/)
func GetHandleFunc(w http.ResponseWriter, r *http.Request) {

	// Allow only GET requests
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Make sure the URL is exactly "/"
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// Render the main HTML page
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
}

// Handles POST requests sent to /ascii-art
func PostHandleFunc(w http.ResponseWriter, r *http.Request) {

	// Allow only POST requests
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Check if the correct route is used
	if r.URL.Path != "/ascii-art" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Get form values from the HTML form
	readText := r.FormValue("text")
	readBanner := r.FormValue("banner")

	// Ensure text and banner are provided
	if readText == "" || readBanner == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest)+": Banner and text required", http.StatusBadRequest)
		return
	}

	// Convert Windows line breaks to Unix line breaks
	readTexts := strings.ReplaceAll(readText, "\r\n", "\n")

	// Generate ASCII art
	result, err := Wrap(readTexts, readBanner)
	if err != nil {
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}

	// Create data to send to the template
	Result := pageData{
		Result: result,
		Text:   readTexts,
		Banner: readBanner,
	}

	// Display result page
	tmpl.ExecuteTemplate(w, "result.html", Result)
}

// Handles banner switching using query parameters
func SwitchHandler(w http.ResponseWriter, r *http.Request) {

	// Allow only GET requests
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Make sure the route is correct
	if r.URL.Path != "/ascii-switch" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// Get text from URL query
	text := r.URL.Query().Get("text")

	// Get banner from URL query
	banner := r.URL.Query().Get("banner")

	// Ensure both values exist
	if text == "" || banner == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest)+": Banner and text requires", http.StatusBadRequest)
		return
	}

	// Generate ASCII art with the selected banner
	result, err := Wrap(text, banner)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Store data for the template
	Result := pageData{
		Text:   text,
		Banner: banner,
		Result: result,
	}

	// Show the result page
	tmpl.ExecuteTemplate(w, "result.html", Result)
}
