package server

import (
	"log"
	"net/http"

	"github.com/doc4suresh/webPageAnalyzer/internal/service"
	"github.com/doc4suresh/webPageAnalyzer/internal/util"
	"github.com/gin-gonic/gin"
)

// analyzeHandler handles GET requests to the /analyze endpoint.
// It expects a 'url' query parameter and returns a JSON response with the analysis results.
// If the 'url' parameter is missing or invalid, it responds with HTTP 400 and an error message.
func analyzeHandler(ctx *gin.Context) {
	// Extract the 'url' query parameter
	url := ctx.Query("url")

	if url == "" {
		log.Print("Missing 'url' query parameter")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Missing 'url' query parameter"})
		return
	}

	if !util.ValidateURL(url) {
		log.Print("Invalid URL format")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid URL format\n" +
			"Please provide a valid URL like\n" +
			"https://www.example.com\n" +
			"or\n" +
			"http://www.example.com"})
		return
	}

	// Call the scraper service
	info, err := service.Scrape(url)
	if err != nil {
		log.Printf("Error analyzing URL %s: %v", url, err)

		// Handle different types of errors with appropriate HTTP status codes
		if contains(err.Error(), "access forbidden") || contains(err.Error(), "403") {
			ctx.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		} else if contains(err.Error(), "too many requests") || contains(err.Error(), "429") {
			ctx.JSON(http.StatusTooManyRequests, gin.H{"message": err.Error()})
		} else if contains(err.Error(), "not found") || contains(err.Error(), "404") {
			ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		} else if contains(err.Error(), "server error") || contains(err.Error(), "500") {
			ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		} else if contains(err.Error(), "connection failed") {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to analyze web page"})
		}
		return
	}

	// Return successful analysis results
	ctx.JSON(http.StatusOK, info)
}

// contains checks if a string contains a substring
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
