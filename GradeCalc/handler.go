package main

import (
	"net/http"
	"html/template"
	"strconv"
)
type PageData struct{
	Math string
	English string
	Science string
	Name string
	Average float64
	Grade string
	Message string
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/"{
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
		load, err := template.ParseFiles("templates/index.html")

	if r.Method == http.MethodGet{
	if err != nil {
		http.Error(w, "Problem Parsing Templates", http.StatusNotFound)
		return 
	}
	load.Execute(w, nil)
}

	if r.Method == http.MethodPost{
    readName := r.FormValue("Name")
	english := r.FormValue("English")
	math := r.FormValue("Math")
	science := r.FormValue("Science")

	if readName == "" || english == "" || math == "" || science == "" {
		http.Error(w, "one of the fields is empty", http.StatusNotFound)
		return
	}

   convert1, err := strconv.ParseFloat(english, 64)
   if err != nil {
	http.Error(w, "Unable to convert", http.StatusNotFound)
	return
   }
   convert2, err := strconv.ParseFloat(math, 64)
   if err != nil {
   http.Error(w, "Unable to convert", http.StatusNotFound)
   return
   }
   convert3, err := strconv.ParseFloat(science, 64)
   if err != nil {
	http.Error(w, "Unable to convert", http.StatusNotFound)
	return
   }
   generate := Average([]float64{convert1, convert2, convert3})
   all := Grade(generate)
 
  // load, err := template.ParseFiles("templates/index.html")
//    if err != nil {
// 	http.Error(w, "error", http.StatusNotFound)
// 	return
//    }
   message := ""
   if all == "A" {
	message = "Excellent"
   } else if all == "B" {
	message = "Good"
   } else if all == "C" {
	message = "Pass"
   } else if all == "D" {
	message = "Poor"
   } else if all == "F" {
	message = "Fail"
   }
   
   result := PageData{
	Math: math,
	English: english, 
	Science: science,
      Name: readName,
	  
	  Average: generate,
	  Grade: all,
	  Message: message,
   }
   load.ExecuteTemplate(w, "index.html", result)
	}
}
