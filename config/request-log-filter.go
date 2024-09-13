package config

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func LogHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s | %s %s | >> %s \n", r.RemoteAddr, r.Method, r.Proto, r.URL)
		fn(w, r)
	}
}

func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}
