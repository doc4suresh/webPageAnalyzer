package server

import (
	"log"
	"net/http"

	scaper "github.com/doc4suresh/webPageAnalyzer/internal/service"
	util "github.com/doc4suresh/webPageAnalyzer/internal/util"
	"github.com/gin-gonic/gin"
)

// analyzeHandler handles GET requests to the /analyze endpoint.
// It expects a 'url' query parameter and returns a JSON response with the URL.
// If the 'url' parameter is missing, it responds with HTTP 400 and an error message.
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

	scaper.Scrape(ctx, url)
}
