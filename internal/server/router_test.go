package server_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Sparker0i/Cactro-Backend-09-Feb-25/internal/config"
	"github.com/Sparker0i/Cactro-Backend-09-Feb-25/internal/server"
)

func TestRouter_PostGetDelete(t *testing.T) {
	// Use a small cache to test capacity-related behavior.
	cfg := &config.Config{MaxCacheSize: 2}
	router := server.NewRouter(cfg)

	// ----- Test POST /cache -----
	recorder := httptest.NewRecorder()
	postBody := bytes.NewBufferString(`{"key": "foo", "value": "bar"}`)
	req, _ := http.NewRequest("POST", "/cache", postBody)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("expected POST status 200, got %d", recorder.Code)
	}

	// ----- Test GET /cache/foo -----
	recorder = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/cache/foo", nil)
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("expected GET status 200, got %d", recorder.Code)
	}
	body, _ := io.ReadAll(recorder.Body)
	if !strings.Contains(string(body), "bar") {
		t.Errorf("expected response body to contain 'bar', got %s", string(body))
	}

	// ----- Test DELETE /cache/foo -----
	recorder = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/cache/foo", nil)
	router.ServeHTTP(recorder, req)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected DELETE status 200, got %d", recorder.Code)
	}

	// Ensure the key is deleted.
	recorder = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/cache/foo", nil)
	router.ServeHTTP(recorder, req)
	if recorder.Code != http.StatusNotFound {
		t.Errorf("expected GET status 404 after deletion, got %d", recorder.Code)
	}
}

func TestRouter_FullCache(t *testing.T) {
	// Create a cache that can hold only 1 item.
	cfg := &config.Config{MaxCacheSize: 1}
	router := server.NewRouter(cfg)

	// Insert the first key.
	recorder := httptest.NewRecorder()
	postBody := bytes.NewBufferString(`{"key": "foo", "value": "bar"}`)
	req, _ := http.NewRequest("POST", "/cache", postBody)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(recorder, req)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected POST status 200 for first key, got %d", recorder.Code)
	}

	// Attempt to insert a second key.
	recorder = httptest.NewRecorder()
	postBody = bytes.NewBufferString(`{"key": "baz", "value": "qux"}`)
	req, _ = http.NewRequest("POST", "/cache", postBody)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(recorder, req)
	if recorder.Code != http.StatusBadRequest {
		t.Errorf("expected POST status 400 when cache is full, got %d", recorder.Code)
	}
}
