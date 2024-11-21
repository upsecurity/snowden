package config

import (
	"log"
	"net/http"
)

func LogHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s | %s %s | >> %s \n", r.RemoteAddr, r.Method, r.Proto, r.URL)
		fn(w, r)
	}
}
