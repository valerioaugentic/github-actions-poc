package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestEnvironmentHandler_Success(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/environment", nil)
	w := httptest.NewRecorder()

	EnvironmentHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response EnvironmentResponse
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response.Environment != "development" {
		t.Errorf("expected environment 'development', got '%s'", response.Environment)
	}

	if response.Version != "SNAPSHOT" {
		t.Errorf("expected version 'SNAPSHOT', got '%s'", response.Version)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type 'application/json', got '%s'", contentType)
	}
}

func TestEnvironmentHandler_WithEnvVars(t *testing.T) {
	os.Setenv("ENVIRONMENT", "production")
	os.Setenv("VERSION", "1.0.0")
	defer func() {
		os.Unsetenv("ENVIRONMENT")
		os.Unsetenv("VERSION")
	}()

	req := httptest.NewRequest(http.MethodGet, "/environment", nil)
	w := httptest.NewRecorder()

	EnvironmentHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response EnvironmentResponse
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response.Environment != "production" {
		t.Errorf("expected environment 'production', got '%s'", response.Environment)
	}

	if response.Version != "1.0.0" {
		t.Errorf("expected version '1.0.0', got '%s'", response.Version)
	}
}

func TestEnvironmentHandler_MethodNotAllowed(t *testing.T) {
	methods := []string{http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch}

	for _, method := range methods {
		t.Run(method, func(t *testing.T) {
			req := httptest.NewRequest(method, "/environment", nil)
			w := httptest.NewRecorder()

			EnvironmentHandler(w, req)

			if w.Code != http.StatusMethodNotAllowed {
				t.Errorf("expected status %d for method %s, got %d", http.StatusMethodNotAllowed, method, w.Code)
			}
		})
	}
}
