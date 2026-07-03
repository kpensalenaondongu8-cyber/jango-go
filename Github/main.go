package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Github struct {
	Avatar    string `json:"avatar_url"`
	Bio       string `json:"bio"`
	RepoCount int    `json:"public_repos"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("Server Running")

	http.ListenAndServe(":8080", nil)
}
func getGit(username string) (Github, error) {
	w, err := http.Get("https://api.github.com/users/" + username)
	if err != nil {
		return Github{}, err
	}
	defer w.Body.Close()
	var j Github

	err = json.NewDecoder(w.Body).Decode(&j)
	return j, err
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("username")

	if text == "" {
		http.Error(w, "empty field", http.StatusBadRequest)
		return
	}
	f, err := getGit(text)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, f)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
