package main

import (
	"fmt"
	"net/http"
)
func main() {
	http.HandleFunc("/hello", HelloFunc)
	http.ListenAndServe(":8080", nil)
}
func HelloFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Hello, World!")
	} else if r.Method == "POST" {
	    name := r.FormValue("name")
		if name == "" {
			fmt.Fprintf(w, "Enter Name")
			return
		} 
	fmt.Fprintf(w, "Hello, %v\n", name)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Status Not Allowed")
	}
}