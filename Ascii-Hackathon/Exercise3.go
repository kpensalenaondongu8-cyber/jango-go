package main

import (
	"net/http"
	"fmt"
	"io"
)

func ma()  {
	http.HandleFunc("/count", CountHandler)
	fmt.Println("Live")
	http.ListenAndServe(":8080", nil)
}

func CountHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Send a POST request with text to count words")
	}else if r.Method == "POST" {
		variable, err := io.ReadAll(r.Body)
		if err != nil{
			http.Error(w, "failed to read the body", http.StatusNotFound)
			return
		}
		vari := len(string(variable))
	fmt.Fprint(w, vari)
	}
}
