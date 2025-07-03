package service

import "testing"

const testURL = "https://www.example.com"

func TestCheckConnection(t *testing.T) {
	err := CheckConnection(testURL)

	if err != nil {
		t.Errorf("CheckConnection(%q) returned error: %v; want no error", testURL, err)
	}
}

func TestGetWebPageInfo(t *testing.T) {
	info, err := GetWebPageInfo(testURL)

	if err != nil {
		t.Errorf("GetWebPageInfo(%q) returned error: %v; want no error", testURL, err)
	}

	if info.Title == "" {
		t.Errorf("GetWebPageInfo(%q) returned empty title; want non-empty title", testURL)
	}
}
