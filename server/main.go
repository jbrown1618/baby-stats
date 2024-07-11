package main

import (
	"fmt"
	"log"
	"net/http"

	"jbrown1618/baby-stats/database"
	"jbrown1618/baby-stats/handler"
	"jbrown1618/baby-stats/middleware"
	"jbrown1618/baby-stats/settings"

	"github.com/gorilla/mux"
)

func main() {
	s := settings.NewApplicationSettings()

	db, err := database.NewApplicationDatabase()
	if err != nil {
		log.Panicln(fmt.Errorf("error getting database: %w", err).Error())
	}
	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/babies", handler.ListBabiesHandler(db)).Methods(http.MethodGet)
	r.HandleFunc("/babies/{babyID:[0-9]+}", handler.GetBabyHandler(db)).Methods(http.MethodGet)
	r.HandleFunc("/babies/{babyID:[0-9]+}/events", handler.ListEventsHandler(db)).Methods(http.MethodGet)
	r.HandleFunc("/babies/{babyID:[0-9]+}/events", handler.CreateEventHandler(db)).Methods(http.MethodPost)

	r.Use(middleware.Logger)
	r.Use(middleware.CommonHeaders(s.IsDev()))

	port := s.ServerPort()
	log.Println("Listening on port " + port)
	http.ListenAndServe(":"+port, r)
}
