package main

import (
	"log"
	"net/http"
)

var availableHandlers []http.HandlerFunc

func createHandler(code int, body string, headers map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("debug: %s %s", r.Method, r.RequestURI)

		for k, v := range headers {
			log.Printf("debug: adding %v, %v", k, v)
			w.Header().Add(k, v)
		}
		w.WriteHeader(code)
		w.Write([]byte(body))
	}
}
