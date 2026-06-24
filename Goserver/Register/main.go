package main

import (
	"fmt"
	"net/http"
)
func main() {
	http.HandleFunc("/register", calculate)
	http.ListenAndServe(":8080", nil)
}
func calculate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
	number := r.FormValue("number")
	if number == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Please provide a number")
	} else if number != "" {
		fmt.Fprintf(w, "You sent the number:%v\n", number)
	}
}
}
