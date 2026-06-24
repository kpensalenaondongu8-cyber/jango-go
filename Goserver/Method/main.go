package main

import (
	"fmt"
	"net/http"
)
func main() {
	http.HandleFunc("/login", LoginFunc)
	http.ListenAndServe(":8080", nil)
}
func LoginFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Please submit your login form")
	} else if r.Method == "POST" {
		fmt.Fprintf(w, "Login successful!")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "405 Method Not Allowed")
	}
}