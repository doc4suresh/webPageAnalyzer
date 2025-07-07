package util

import "testing"

func TestValidateURL(t *testing.T) {
	tests := []struct {
		name     string
		rawurl   string
		expected bool
	}{
		{"Valid URL", "https://www.example.com", true},
		{"Invalid URL", "example", false},
		{"Empty URL", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateURL(tt.rawurl)
			if result != tt.expected {
				t.Errorf("ValidateURL(%q) = %v; want %v", tt.rawurl, result, tt.expected)
			}
		})
	}
}
