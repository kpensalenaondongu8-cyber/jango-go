package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Joke struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
	Rating    string `json:"rating"`
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("server running")

	http.ListenAndServe(":8080", nil)
}

func getJoke() (Joke, error) {
	w, err := http.Get("https://official-joke-api.appspot.com/random_joke")
	if err != nil {
		return Joke{}, err
	}
	defer w.Body.Close()
	var j Joke
	err = json.NewDecoder(w.Body).Decode(&j)
	return j, err
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	q, err := getJoke()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, q)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
