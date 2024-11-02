package main

import (
	"encoding/json"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Hello, World!",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func setupRoutes() {
	http.HandleFunc("/", helloHandler)
}

func main() {
	setupRoutes()

	http.ListenAndServe(":8080", nil)
}
