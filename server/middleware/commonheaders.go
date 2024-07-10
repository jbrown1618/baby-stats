package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
)

func CommonHeaders(isDev bool) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			if isDev {
				w.Header().Set("Access-Control-Allow-Origin", "*")
			}
			next.ServeHTTP(w, r)
		})
	}
}
