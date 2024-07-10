package handler

import (
	"encoding/json"
	"fmt"
	"jbrown1618/baby-stats/database"
	"log"
	"net/http"
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
