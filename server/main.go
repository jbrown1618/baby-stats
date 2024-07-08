package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"jbrown1618/baby-stats/database"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(fmt.Errorf("error loading environment variables: %w", err).Error())
	}

	port := os.Getenv("PORT")
	isDev := os.Getenv("ENVIRONMENT_TYPE") == "DEV"

	db, err := database.GetApplicationDatabase()
	if err != nil {
		fmt.Println(fmt.Errorf("error getting database: %w", err).Error())
		os.Exit(1)
	}
	defer db.Close()

	http.HandleFunc("/babies", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request from the mobile app")
		w.Header().Set("Content-Type", "application/json")
		if isDev {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		babies, err := database.ListBabies(db, 1)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - something went wrong"))
			return
		}

		babiesJson, err := json.Marshal(babies)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - something went wrong"))
			return
		}
		fmt.Fprint(w, string(babiesJson))
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Listening on port " + port)
	http.ListenAndServe(":"+port, nil)
}
