package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type WeatherResponse struct {
	CurrentWeather struct {
		Temperature   float64 `json:"temperature"`
		Windspeed     float64 `json:"windspeed"`
		Time          string  `json:"time"`
		Interval      int     `json:"interval"`
		Weathercode   int     `json:"weathercode"`
		Winddirection int     `json:"winddirection"`
	} `json:"current_weather"`
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("Server Running")

	http.ListenAndServe(":8080", nil)
}
func getWeather(latitude string, longitude string) (WeatherResponse, error) {

	w, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=" + latitude + "&longitude=" + longitude + "&current_weather=true")
	if err != nil {
		return WeatherResponse{}, err
	}
	defer w.Body.Close()
	var j WeatherResponse
	err = json.NewDecoder(w.Body).Decode(&j)
	return j, err
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fetchLatt := r.FormValue("latitude")
	fetchLong := r.FormValue("longitude")
	if fetchLatt == "" || fetchLong == "" {
		http.Error(w, "country isn't found", http.StatusNotFound)
		return
	}
	e, err := getWeather(fetchLatt, fetchLong)
	if err != nil {
		http.Error(w, "internla server error", http.StatusInternalServerError)
		return
	}
	j := WeatherResponse{
		CurrentWeather: e.CurrentWeather,
	}
	err = tmpl.Execute(w, j)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

}
