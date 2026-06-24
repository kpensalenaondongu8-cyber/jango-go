package main

import (
	"fmt"
	"net/http"
)
func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)
	http.HandleFunc("/contact", ContactHandler)
	http.ListenAndServe(":8080", nil)
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page")
}
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact Page")
}