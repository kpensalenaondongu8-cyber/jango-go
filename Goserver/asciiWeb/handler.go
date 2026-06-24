package main

import (
	"html/template"
	"net/http"
	"strings"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

type pageData struct {
	Result string
	Text   string
	Banner string
}

func GetHandleFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
}

func PostHandleFunc(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/ascii-art" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	readText := r.FormValue("text")
	readBanner := r.FormValue("banner")

	if readText == "" || readBanner == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	readTexts := strings.ReplaceAll(readText, "\r\n", "\n")

	q, err := Wrap(readTexts, readBanner)
	if err != nil {
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	Result := pageData{
		Result: q,
		Text:   readTexts,
		Banner: readBanner,
	}
	tmpl.ExecuteTemplate(w, "m.html", Result)
}
func SwitchHandler(w http.ResponseWriter, r *http.Request) {

	text := r.URL.Query().Get("text")

	banner := r.URL.Query().Get("banner")

	x, err := Wrap(text, banner)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	Result := pageData{
		Text:   text,
		Banner: banner,
		Result: x,
	}
	tmpl.ExecuteTemplate(w, "m.html", Result)
}
