package api

import (
	"encoding/json"
	"net/http"
	"snowden/config"

	"github.com/gorilla/mux"
)

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/", config.LogHandler(makeHTTPHandleFunc(SlashHandler))).Methods("GET")

	router.HandleFunc("/api/v1/health", config.LogHandler(makeHTTPHandleFunc(HealthHandler))).Methods("GET")
	router.HandleFunc("/api/v1/vulnerability/cve", config.LogHandler(makeHTTPHandleFunc(ReadVulnerabilityByCve))).Methods("GET")
	// router.HandleFunc("/api/v1/vulnerability/cwe", config.LogHandler(makeHTTPHandleFunc(ReadVulnerabilityByCwe))).Methods("GET")

	err := http.ListenAndServe(s.Port, router)
	if err != nil {
		return
	}
}

func SlashHandler(w http.ResponseWriter, r *http.Request) error {
	return WriteJson(w, http.StatusOK, API{message: "Welcome to Snowden API"})
}

func NewApiServer(port string) *APIServer {
	return &APIServer{
		Port: port,
	}
}

func WriteJson(w http.ResponseWriter, status int, value any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(value)
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			err := WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
			if err != nil {
				return
			}
		}
	}
}

type APIServer struct {
	Port string
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

type API struct {
	message string
}
