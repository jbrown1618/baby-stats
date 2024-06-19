package main

import (
	"fmt"
	"net/http"
)

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request from the mobile app")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
        fmt.Fprintf(w, "{ \"message\": \"Welcome to my website!\" }")
    })

    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Listening on port 8082")
    http.ListenAndServe(":8082", nil)
}