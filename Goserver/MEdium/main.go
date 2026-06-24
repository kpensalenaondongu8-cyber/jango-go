package main

import (
	"fmt"
	"net/http"
)
func main() {
	http.HandleFunc("/profile", ProfileFunc) 
	http.HandleFunc("/update", UpdateFunc)
	http.ListenAndServe(":8080", nil)
}
func ProfileFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method !=  "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method Not allowed")
		return
	}
	username := r.URL.Query().Get("username")
	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Please Provide a Username")
		return
	} else if username != "" {
		fmt.Fprintf(w, "Profile page of:%v\n", username)
	} 
}
func UpdateFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Please use Post to update")
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method Not Allowed")
		return
	}
	Username := r.FormValue("username")
	bio := r.FormValue("bio")
	if Username != "" && bio != "" {
		fmt.Fprintf(w, "Profile updated! Usename: %v, Bio: %v\n", Username, bio)
	} else {	
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Both username and bio are required")
	}
}