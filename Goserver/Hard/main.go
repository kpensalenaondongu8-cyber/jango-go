package main

import (
	"fmt"
	"net/http"
)
func main() {
	http.HandleFunc("/register", RegisterFunc)
	http.ListenAndServe(":8080", nil)
}
func RegisterFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Please submit the registeration form")
	} else if r.Method == "POST" {
	username :=	r.FormValue("username")
	email := r.FormValue("email")
	if username != "" && email != "" {
		fmt.Fprintf(w, "Welcome %v! Your email is %v", username, email)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Username and email are required")
	}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Mthod Not Allowed")
	}
}