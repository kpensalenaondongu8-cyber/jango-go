package main

import (
	"fmt"
	"net/http"
)
func main() {
	http.HandleFunc("/", GetHandler)
	//http.HandleFunc("/result", PostHandler)

	fmt.Println("Server Running on Port... ")
	http.ListenAndServe(":8080", nil)
}