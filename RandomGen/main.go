package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Random struct {
	Results []struct {
		Name struct {
			Title string `json:"title"`
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Location struct {
			Name    string `json:"name"`
			City    string `json:"city"`
			State   string `json:"state"`
			Country string `json:"country"`
		} `json:"location"`
		Email   string `json:"email"`
		Picture struct {
			Large     string `json:"large"`
			Medium    string `json:"medium"`
			Thumbnail string `json:"thumbnail"`
		} `json:"picture"`
	} `json:"results"`
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("server runnning")

	http.ListenAndServe(":8080", nil)
}
func rand() (Random, error) {
	w, err := http.Get("https://randomuser.me/api/")
	if err != nil {
		return Random{}, err
	}
	defer w.Body.Close()

	var j Random
	json.NewDecoder(w.Body).Decode(&j)
	return j, err
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	x, err := rand()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, x)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
