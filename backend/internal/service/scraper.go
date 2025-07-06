package service

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gocolly/colly"
)

// Scrape performs web page analysis and returns the results
func Scrape(url string) (*WebPageInfo, error) {
	if err := validateURL(url); err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}

	info, err := analyzeWebPage(url)
	if err != nil {
		return nil, fmt.Errorf("analysis failed: %w", err)
	}

	return info, nil
}

// validates the URL and checks if it's reachable
func validateURL(url string) error {
	collector := createCollector()

	retryConfig := getRetryConfig()
	var capturedError error

	// Set up error handling
	collector.OnError(func(r *colly.Response, err error) {
		capturedError = createSpecificError(r.StatusCode, r.Request.URL.String(), err)
		log.Printf("HTTP Error %d for %s: %v", r.StatusCode, r.Request.URL, err)
	})

	// Attempt to reach the URL with retries
	for attempt := 1; attempt <= retryConfig.limit; attempt++ {
		err := collector.Visit(url)

		// Return specific error if captured (403, 429, etc.)
		if capturedError != nil {
			return capturedError
		}

		if err != nil {
			log.Printf("Attempt %d: Failed to reach %s, error: %v", attempt, url, err)
			if attempt < retryConfig.limit {
				time.Sleep(time.Duration(retryConfig.delay) * time.Second)
			}
		} else {
			log.Printf("Successfully reached %s on attempt %d", url, attempt)
			return nil
		}
	}

	return fmt.Errorf("failed to reach %s after %d attempts", url, retryConfig.limit)
}

// analyzeWebPage performs the actual web page analysis
func analyzeWebPage(url string) (*WebPageInfo, error) {
	collector := createCollector()

	info := &WebPageInfo{
		URL:               url,
		HeadCount:         make(map[string]int),
		AccessibleLinks:   0,
		InAccessibleLinks: 0,
		IsLoginForm:       false,
	}

	// Set up HTML element handlers
	setupHTMLHandlers(collector, info)

	if err := collector.Visit(url); err != nil {
		return nil, fmt.Errorf("failed to visit %s: %w", url, err)
	}

	return info, nil
}

// createCollector creates and configures a new Colly collector
func createCollector() *colly.Collector {
	collector := colly.NewCollector()

	// Set realistic user agent to avoid being blocked
	collector.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"

	// Add request headers to appear more like a real browser
	collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.5")
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
	})

	return collector
}

// setupHTMLHandlers configures all HTML element handlers for data extraction
func setupHTMLHandlers(collector *colly.Collector, info *WebPageInfo) {
	// Extract HTML version
	collector.OnHTML("html", func(e *colly.HTMLElement) {
		if doctype := e.Attr("html"); doctype != "" {
			info.HTMLVersion = doctype
		}
	})

	// Extract page title
	collector.OnHTML("title", func(e *colly.HTMLElement) {
		info.Title = e.Text
	})

	// Count heading tags
	collector.OnHTML("h1, h2, h3, h4, h5, h6", func(e *colly.HTMLElement) {
		info.HeadCount[e.Name]++
	})

	// Extract accessible links
	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if href != "" && href != "#" {
			info.AccessibleLinks++
		}
	})

	// Extract inaccessible links (JavaScript, mailto, etc.)
	collector.OnHTML("a[href^='javascript:'], a[href^='mailto:'], a[href^='tel:']", func(e *colly.HTMLElement) {
		info.InAccessibleLinks++
	})

	// Detect login forms
	collector.OnHTML("form", func(e *colly.HTMLElement) {
		action := e.Attr("action")
		id := e.Attr("id")
		class := e.Attr("class")

		// Check for common login form indicators
		if containsLoginIndicator(action) || containsLoginIndicator(id) || containsLoginIndicator(class) {
			info.IsLoginForm = true
		}
	})
}

// createSpecificError creates a specific error message based on HTTP status code
func createSpecificError(statusCode int, url string, err error) error {
	switch statusCode {
	case 403:
		return fmt.Errorf("access forbidden (403) for %s - website may block automated requests", url)
	case 429:
		return fmt.Errorf("too many requests (429) for %s - rate limited", url)
	case 404:
		return fmt.Errorf("page not found (404) for %s", url)
	case 500:
		return fmt.Errorf("server error (500) for %s", url)
	default:
		return fmt.Errorf("HTTP error %d for %s: %v", statusCode, url, err)
	}
}

// getRetryConfig reads retry configuration from environment variables
func getRetryConfig() struct {
	limit int
	delay int
} {
	retryDelayStr := os.Getenv("URL_RETRY_DELAY")
	retryDelay, err := strconv.Atoi(retryDelayStr)
	if err != nil || retryDelay < 1 {
		retryDelay = 1
	}

	retryLimitStr := os.Getenv("URL_RETRY_LIMIT")
	retryLimit, err := strconv.Atoi(retryLimitStr)
	if err != nil || retryLimit < 1 {
		retryLimit = 3
	}

	return struct {
		limit int
		delay int
	}{
		limit: retryLimit,
		delay: retryDelay,
	}
}

// containsLoginIndicator checks if a string contains login-related keywords
func containsLoginIndicator(text string) bool {
	loginKeywords := []string{"login", "signin", "auth", "authenticate"}
	text = string(text)
	for _, keyword := range loginKeywords {
		if contains(text, keyword) {
			return true
		}
	}
	return false
}

// contains checks if a string contains a substring (case-insensitive)
func contains(s, substr string) bool {
	return len(s) >= len(substr) &&
		(s == substr ||
			(len(s) > len(substr) &&
				(s[:len(substr)] == substr ||
					s[len(s)-len(substr):] == substr ||
					containsSubstring(s, substr))))
}

// containsSubstring is a simple substring check
func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
