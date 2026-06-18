package main

import (
	"net/http"
	"fmt"
)

func ain()  {
	http.HandleFunc("/ping", pongHandler)
	fmt.Println("Live")
	http.ListenAndServe(":8080", nil)
}

func pongHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "pong")
}
