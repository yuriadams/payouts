package controllers

import (
	"fmt"
	"log"
	"net/http"
)

type Headers map[string]string

func respondWith(w http.ResponseWriter, status int, headers Headers) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)
}

func RespondWithJSON(w http.ResponseWriter, status int, response string) {
	respondWith(w, status, Headers{"Content-Type": "application/json"})
	fmt.Fprintf(w, response)
}

func Logging(format string, v ...interface{}) {
	log.Printf(fmt.Sprintf("%s\n", format), v...)
}
