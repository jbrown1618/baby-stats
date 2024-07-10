package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"jbrown1618/baby-stats/database"
	"jbrown1618/baby-stats/handler"
	"jbrown1618/baby-stats/middleware"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(fmt.Errorf("error loading environment variables: %w", err).Error())
	}

	port := os.Getenv("PORT")
	isDev := os.Getenv("ENVIRONMENT_TYPE") == "DEV"

	db, err := database.NewApplicationDatabase()
	if err != nil {
		log.Println(fmt.Errorf("error getting database: %w", err).Error())
		os.Exit(1)
	}
	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/babies", handler.BabiesHandler(db)).Methods(http.MethodGet)
	r.HandleFunc("/babies/{babyID:[0-9]+}", handler.BabyHandler(db)).Methods(http.MethodGet)
	r.HandleFunc("/babies/{babyID:[0-9]+}/events", handler.EventsHandler(db)).Methods(http.MethodGet)

	r.Use(middleware.Logger)
	r.Use(middleware.CommonHeaders(isDev))

	log.Println("Listening on port " + port)
	http.ListenAndServe(":"+port, r)
}
