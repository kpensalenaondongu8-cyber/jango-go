package main

import (
	"net/http"
	"fmt"
)

func main()  {
	http.HandleFunc("/legacy", LegacydHandler)
	http.HandleFunc("/v2")
	fmt.Println("Live")
	http.ListenAndServe(":8080", nil)
}
func LegacydHandler(w http.ResponseWriter, r *http.Request) {
	
	
}


