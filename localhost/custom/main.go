package main

import (
	"fmt"
	"net/http"
)
func main() {
	http.HandleFunc("/greet", usersHandlefunc)
	http.ListenAndServe(":8080", http.NewServeMux())
}
func usersHandlefunc(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
    if name == "" {
		fmt.Fprintf(w, "Hello, Stranger")
	} else if name != "" {
		fmt.Fprintf(w, "Hello %v", name)
	}
}