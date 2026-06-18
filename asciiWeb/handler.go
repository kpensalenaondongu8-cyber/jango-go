package main

import (
	"net/http"
	"html/template"
)

type PageData struct {
	Result string
}
func GetHandleFunc(w http.ResponseWriter, r *http.Request){
	load, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "could'nt parse files", http.StatusNotFound)
		return
	}
	load.Execute(w, nil)
}
func PostHandleFunc(w http.ResponseWriter, r *http.Request) {
	readText := r.FormValue("text")
	readBanner := r.FormValue("banner")

	if readText == ""{
		http.Error(w, "No Text Entered", http.StatusNotFound)
		return
	}
	if readBanner == ""{
		http.Error(w, "No Banner File Selected", http.StatusNotFound)
		return
	}

	generate, err := Wrap(readText, readBanner)
	if err != nil {
		http.Error(w, "could'nt read", http.StatusNotFound)
		return
	}
	result := PageData{
		Result: generate,
	}
	load, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "could'nt read", http.StatusNotFound)
		return
	}
	load.Execute(w, result)
}