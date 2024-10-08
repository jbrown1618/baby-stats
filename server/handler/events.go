package handler

import (
	"encoding/json"
	"fmt"
	"jbrown1618/baby-stats/database"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ListEventsHandler(db *database.ApplicationDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		babyIDPath := mux.Vars(r)["babyID"]
		babyID, err := strconv.ParseUint(babyIDPath, 10, 64)
		if err != nil {
			log.Printf("invalid baby ID %s: %s", babyIDPath, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - something went wrong"))
			return
		}

		events, err := db.ListEvents(1, babyID)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - something went wrong"))
			return
		}

		eventsJson, err := json.Marshal(events)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - something went wrong"))
			return
		}

		fmt.Fprint(w, string(eventsJson))
	}
}

func CreateEventHandler(db *database.ApplicationDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		babyIDPath := mux.Vars(r)["babyID"]
		babyID, err := strconv.ParseUint(babyIDPath, 10, 64)
		if err != nil {
			log.Printf("invalid baby ID %s: %s", babyIDPath, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - something went wrong"))
			return
		}

		var event database.Event
		jsonDecoder := json.NewDecoder(r.Body)
		err = jsonDecoder.Decode(&event)
		if err != nil {
			log.Printf("invalid request body: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - something went wrong"))
			return
		}

		id, err := db.CreateEvent(babyID, &event)
		if err != nil {
			log.Printf("error creating event: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - something went wrong"))
			return
		}

		fmt.Fprintf(w, `{ "id": %d }`, id)
	}
}
