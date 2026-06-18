package main

import (
	"html/template"
	"net/http"
)
func main() {
	http.HandleFunc("/greet", handlerFunc)
	http.ListenAndServe(":8080", nil)
}
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
   x, err := template.ParseFiles("template/index.html")
   if err != nil {
	http.Error(w, "Template failed to load", 404)
	return
   }
   if name != "" {
   x.Execute(w, name)
   } else {
	x.Execute(w, "Hello stranger!")
   }
}