package main

import (
	"fmt"
	"net/http"
)
func main() {
	http.HandleFunc("/", usersHandlefunc)
	//http.HandleFunc("/users", usersHandlefunc)
	http.ListenAndServe(":8080", nil)
}
func usersHandlefunc(w http.ResponseWriter, r *http.Request){
	fmt.Println("Hello over there")
	fmt.Fprintf(w, "Hi, thanks for calling my /user API with HTTP method '%v'", r.Method)
}
