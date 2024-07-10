package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

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

	db, err := database.NewApplicationDatabase()
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
		babies, err := db.ListBabies(1)
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

	http.HandleFunc("/babies/{babyID}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request from the mobile app")
		w.Header().Set("Content-Type", "application/json")
		if isDev {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		babyIDPath := r.PathValue("babyID")
		babyID, err := strconv.ParseUint(babyIDPath, 10, 64)
		if err != nil {
			fmt.Printf("invalid baby ID %s: %s", babyIDPath, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - something went wrong"))
			return
		}

		baby, err := db.GetBaby(1, babyID)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - something went wrong"))
			return
		}

		babyJson, err := json.Marshal(baby)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - something went wrong"))
			return
		}
		fmt.Fprint(w, string(babyJson))
	})

	http.HandleFunc("/babies/{babyID}/events", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request from the mobile app")
		w.Header().Set("Content-Type", "application/json")
		if isDev {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		babyIDPath := r.PathValue("babyID")
		babyID, err := strconv.ParseUint(babyIDPath, 10, 64)
		if err != nil {
			fmt.Printf("invalid baby ID %s: %s", babyIDPath, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - something went wrong"))
			return
		}

		events, err := db.ListEvents(1, babyID)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - something went wrong"))
			return
		}

		eventsJson, err := json.Marshal(events)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - something went wrong"))
			return
		}
		fmt.Fprint(w, string(eventsJson))
	})

	fmt.Println("Listening on port " + port)
	http.ListenAndServe(":"+port, nil)
}
