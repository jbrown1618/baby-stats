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

func BabiesHandler(db *database.ApplicationDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		babies, err := db.ListBabies(1)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - something went wrong"))
			return
		}

		babiesJson, err := json.Marshal(babies)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - something went wrong"))
			return
		}

		fmt.Fprint(w, string(babiesJson))
	}
}

func BabyHandler(db *database.ApplicationDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		babyIDPath := mux.Vars(r)["babyID"]
		babyID, err := strconv.ParseUint(babyIDPath, 10, 64)
		if err != nil {
			log.Printf("invalid baby ID %s: %s", babyIDPath, err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - invalid baby ID"))
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
			log.Printf("Error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - something went wrong"))
			return
		}

		fmt.Fprint(w, string(babyJson))
	}
}
