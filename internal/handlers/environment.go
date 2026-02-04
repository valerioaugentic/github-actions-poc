package handlers

import (
	"encoding/json"
	"net/http"
	"os"
)

type EnvironmentResponse struct {
	Environment string `json:"environment"`
	Version     string `json:"version"`
	Message     string `json:"message"`
}

func EnvironmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "development"
	}

	version := os.Getenv("VERSION")
	if version == "" {
		version = "SNAPSHOT"
	}

	response := EnvironmentResponse{
		Environment: environment,
		Version:     version,
		Message:     "Hello, World!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
