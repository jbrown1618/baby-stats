package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

type SimpleResponseWriter struct {
	http.ResponseWriter
	buf *bytes.Buffer
}

func (mrw *SimpleResponseWriter) Write(p []byte) (int, error) {
	return mrw.buf.Write(p)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read the body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		log.Printf("%s\t%s\t%s", r.Method, r.RequestURI, body)

		// And now set a new body, which will simulate the same data we read:
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		next.ServeHTTP(w, r)
	})
}
