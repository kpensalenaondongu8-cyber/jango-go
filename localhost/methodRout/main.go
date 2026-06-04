package main

import (
	"fmt"
	"net/http"
)
func usersHandlefunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Send me data via POST")
	} else if r.Method == "POST" {
		fmt.Fprintf(w, "Thanks for the data")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method Not ALlowed")
	}
}
func main() {
	http.HandleFunc("/", usersHandlefunc)
	http.ListenAndServe(":8080", nil)
}