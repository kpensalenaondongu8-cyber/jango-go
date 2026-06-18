package main

import (
	"net/http"
	"fmt"
)

func main()  {
	http.HandleFunc("/dashboard", dashboardHandler)
	fmt.Println("Live")
	http.ListenAndServe(":8080", nil)
}
func dashboardHandler(w http.ResponseWriter, r *http.Request) {

	headervalue := r.Header.Get("X-API-Key")
	if headervalue != "secret123" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Fprint(w, "Welcome")
	
	
}
