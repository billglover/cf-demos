package main

import (
	"net/http"
	"os"
)

var color string

func main() {
	color = os.Getenv("DEPLOYMENT_COLOR")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

// Hello handles HTTP requests and returns the value of the color variable.
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Active deployment: " + color + "\n"))
}
