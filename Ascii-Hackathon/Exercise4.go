package main

import (
	"net/http"
	"fmt"
	"strconv"
)

func bedwt()  {
	http.HandleFunc("/calculate", CalculateHandler)
	fmt.Println("Live")
	http.ListenAndServe(":8080", nil)
}

func CalculateHandler(w http.ResponseWriter, r *http.Request)  {
	Query1 := r.URL.Query().Get("op")
	Query2 := r.URL.Query().Get("a")
	Query3 := r.URL.Query().Get("b")

	num1, err := strconv.Atoi(Query2)
	if err != nil {
		http.Error(w, "error", 400)
		return
	}
	num2, err := strconv.Atoi(Query3)
	if err != nil {
		http.Error(w, "error", 400)
		return
	}
	if Query1 == "add" {
		fmt.Fprint(w, num1 + num2)
		return
	}
	if Query1 == "substract" {
		fmt.Fprint(w, num1 - num2)
		return
	}
	if Query1 == "multiply" {
		fmt.Fprint(w, num1 * num2)
		return
	}
	if Query1 != "add" || Query1 != "substract" || Query1 != "multiply"{
		http.Error(w, "Bad Request", 400)
		return
	}

}
