package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	// listen to port
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Welcome to new server!")
}
