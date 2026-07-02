package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Facts struct {
	Facts string `json:"facts"`
	Count int    `json:"count"`
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("Server Running")

	http.ListenAndServe(":8080", nil)
}
func getFacts() (Facts, error) {
	w, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		return Facts{}, err
	}

	var j Facts
	err = json.NewDecoder(w.Body).Decode(&j)
	w.Body.Close()

	return j, err
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	count, err := strconv.Atoi(text)
	if err != nil {
		count = 0
	}
	count++
	x, err := getFacts()

	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	data := Facts{
		Facts: x.Facts,
		Count: count,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
