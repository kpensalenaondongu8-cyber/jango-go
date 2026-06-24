package main

import (
	"fmt"
	"net/http"
)
func main() {
	http.HandleFunc("/welcome", WelcomeFunc)
	http.HandleFunc("/submit", SubmitFunc)
	http.ListenAndServe(":8080", nil)
}
func WelcomeFunc(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if r.Method == "POST" {
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(w, "POST not allowed")
	}else if name != "" {
		 fmt.Fprintf(w, "Welcome, %v!\nMethod used: GET", name)
	} else {
          fmt.Fprintf(w, "Please provide your name")
	}
}
func SubmitFunc(w http.ResponseWriter, r *http.Request) {
   if r.Method == "GET" {
	fmt.Fprintf(w, "Please use POST to submit")
   } else if r.Method == "POST" {
	fmt.Fprintf(w, "Form submitted successfully")
   }
}
