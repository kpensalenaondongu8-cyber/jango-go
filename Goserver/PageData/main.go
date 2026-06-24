package main

import (
	"html/template"
	"net/http"
)
type PageData struct {
	Name string
	City string
}
func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8080", nil)
}
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	x, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Templates could'nt load", 404)
       return
	}
	word := PageData{
		Name: "Thomas",
		City: "Gboko",
	}
	x.Execute(w, word)
}