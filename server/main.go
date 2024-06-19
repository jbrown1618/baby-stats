package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(fmt.Errorf("error loading environment variables: %w", err).Error())
	}

	port := os.Getenv("PORT")
	isDev := os.Getenv("ENVIRONMENT_TYPE") == "DEV"
	
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request from the mobile app")
		w.Header().Set("Content-Type", "application/json")
		if isDev {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
        fmt.Fprintf(w, "{ \"message\": \"Welcome to my website!\" }")
    })

    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Listening on port " + port)
    http.ListenAndServe(":" + port, nil)
}