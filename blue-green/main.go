package main

import (
	"fmt"
	"net/http"
	"os"
)

var colour string

func main() {
	colour = os.Getenv("DEPLOYMENT_COLOR")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

// Hello handles HTTP requests and returns the value of the color variable.
func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Deployment-Colour", colour)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s!", colour)
}
