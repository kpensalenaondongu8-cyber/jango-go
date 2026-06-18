package main

import (
	"net/http"
	"fmt"
)

func mai()  {
	http.HandleFunc("/hello", HelloHandler)
	fmt.Println("Live")
	http.ListenAndServe(":8080", nil)
}
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintf(w, "Hello, Guest!")
	} else{
        fmt.Fprintf(w, "Hello, %v!", name)
	}
}