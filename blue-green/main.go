package main

import (
	"encoding/json"
	"net/http"
	"os"
)

// Payload is the response returned by the server. It simply indicates whether
// this is a blue or a green deployment.
type Payload struct {
	Color string `json:"color,omitempty"`
}

var color string

func main() {
	color = os.Getenv("DEPLOYMENT_COLOR")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

// Hello handles HTTP requests and returns the value of the color variable.
func hello(w http.ResponseWriter, r *http.Request) {
	p := Payload{Color: color}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
	w.WriteHeader(http.StatusOK)
}
