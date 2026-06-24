package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", GetHandleFunc)
	mux.HandleFunc("/ascii-switch", SwitchHandler)
	mux.HandleFunc("/ascii-art", PostHandleFunc)

	fmt.Println("Server Running.....")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

}
