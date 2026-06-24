package main

import (
	"fmt"
	"net/http"
)
func main() {
	http.HandleFunc("/search", searchFunc)
	http.ListenAndServe(":8080", nil)
}
func searchFunc(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if q != "" {
		fmt.Fprintf(w, "You searched for:%v\n", q)
	} else {
        fmt.Fprintf(w, "Please Provide a search term!")
	}
}