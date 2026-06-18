package main

import (
	"fmt"
	"net/http"
	"html/template"
)
func main() {
	http.HandleFunc("/", handlerfunc)
	http.ListenAndServe(":8080", nil)
}
func handlerfunc(w http.ResponseWriter, r *http.Request) {
	x, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Error(w, "Templates failed to load", 404)
		return
	}
    x.Execute(w, nil)	
}