package service

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestValidateURL(t *testing.T) {
	// Mock HTTP server to simulate a valid URL response
	// This server will return a 200 OK status
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	err := validateURL(ts.URL)

	if err != nil {
		t.Errorf("Test validateURL(%q) unexpected error: %v", ts.URL, err)
	}
}

func TestAnalyzeWebPage(t *testing.T) {
	// Mock HTTP server to simulate a web page response
	// This server will return a simple HTML page with a title "Mock Title"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<html><head><title>Mock Title</title></head><body></body></html>`))
	}))
	defer ts.Close()

	results, err := analyzeWebPage(ts.URL)

	if err != nil {
		t.Errorf("Test analyzeWebPage unexpected error: %v", err)
	}

	if results.Title != "Mock Title" {
		t.Errorf("Test analyzeWebPage expected title 'Mock Title', got %q", results.Title)
	}
}

func TestScrape(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<html><head><title>Scrape Title</title></head><body><h1>Header</h1><form action="/login"></form></body></html>`))
	}))
	defer ts.Close()

	info, err := Scrape(ts.URL)

	if err != nil {
		t.Errorf("Scrape(%q) unexpected error: %v", ts.URL, err)
	}

	if info.Title != "Scrape Title" {
		t.Errorf("Scrape(%q) expected title 'Scrape Title', got %q", ts.URL, info.Title)
	}

	if !info.IsLoginForm {
		t.Errorf("Scrape(%q) expected IsLoginForm true, got false", ts.URL)
	}

	if info.HeadCount["h1"] != 1 {
		t.Errorf("Scrape(%q) expected 1 h1, got %d", ts.URL, info.HeadCount["h1"])
	}
}

func TestCreateSpecificError(t *testing.T) {
	tests := []struct {
		status   int
		url      string
		err      error
		expected string
	}{
		{403, "http://x", nil, "access forbidden (403) for http://x"},
		{429, "http://y", nil, "too many requests (429) for http://y"},
		{404, "http://z", nil, "page not found (404) for http://z"},
		{500, "http://a", nil, "server error (500) for http://a"},
		{418, "http://b", fmt.Errorf("teapot"), "HTTP error 418 for http://b: teapot"},
	}

	for _, tt := range tests {
		err := createSpecificError(tt.status, tt.url, tt.err)
		if err == nil || err.Error()[:len(tt.expected)] != tt.expected {
			t.Errorf("createSpecificError(%d, %q, %v) = %v, want prefix %q", tt.status, tt.url, tt.err, err, tt.expected)
		}
	}
}

func TestGetRetryConfig(t *testing.T) {
	os.Setenv("URL_RETRY_DELAY", "2")
	os.Setenv("URL_RETRY_LIMIT", "4")

	cfg := getRetryConfig()

	if cfg.delay != 2 || cfg.limit != 4 {
		t.Errorf("getRetryConfig() with env = %+v, want delay=2, limit=4", cfg)
	}
}

func TestContainsLoginIndicator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"login", true},
		{"user-login-form", true},
		{"auth", true},
		{"signin", true},
		{"authenticate", true},
		{"register", false},
		{"", false},
	}

	for _, tt := range tests {
		if got := containsLoginIndicator(tt.input); got != tt.expected {
			t.Errorf("containsLoginIndicator(%q) = %v, want %v", tt.input, got, tt.expected)
		}
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		s, substr string
		expected  bool
	}{
		{"hello", "ell", true},
		{"hello", "lo", true},
		{"hello", "he", true},
		{"hello", "world", false},
		{"", "", true},
		{"a", "a", true},
		{"a", "b", false},
	}

	for _, tt := range tests {
		if got := contains(tt.s, tt.substr); got != tt.expected {
			t.Errorf("contains(%q, %q) = %v, want %v", tt.s, tt.substr, got, tt.expected)
		}
	}
}

func TestContainsSubstring(t *testing.T) {
	tests := []struct {
		s, substr string
		expected  bool
	}{
		{"hello", "ell", true},
		{"hello", "lo", true},
		{"hello", "he", true},
		{"hello", "world", false},
		{"", "", true},
		{"a", "a", true},
		{"a", "b", false},
	}

	for _, tt := range tests {
		if got := containsSubstring(tt.s, tt.substr); got != tt.expected {
			t.Errorf("containsSubstring(%q, %q) = %v, want %v", tt.s, tt.substr, got, tt.expected)
		}
	}
}
