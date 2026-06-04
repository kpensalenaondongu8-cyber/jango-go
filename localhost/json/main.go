package main

import(
	"net/http"
	"encoding/json"
)
func main() {
	http.HandleFunc("/greet", usersHandlefunc)
	http.ListenAndServe(":8080", nil)
}
func usersHandlefunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	x := map[string]string {"message": "Hello", "timestamp": "1234567890"}
   json.NewEncoder(w).Encode(x)
}
