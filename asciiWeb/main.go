
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", GetHandleFunc)
    http.HandleFunc("/ascii-art", PostHandleFunc)

    fmt.Println("Server Running.....")
    http.ListenAndServe(":8080", nil)
}