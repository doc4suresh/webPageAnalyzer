package server

import (
	"log"
	"net/http"

	util "github.com/doc4suresh/webPageAnalyzer/internal/Util"
	"github.com/gin-gonic/gin"
)

// analyzeHandler handles GET requests to the /analyze endpoint.
// It expects a 'url' query parameter and returns a JSON response with the URL.
// If the 'url' parameter is missing, it responds with HTTP 400 and an error message.
func analyzeHandler(ctx *gin.Context) {
	// Extract the 'url' query parameter
	url := ctx.Query("url")

	if url == "" {
		log.Fatal("Missing 'url' query parameter")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'url' query parameter"})
		return
	}

	if !util.ValidateURL(url) {
		log.Fatal("Invalid URL format")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format"})
		return
	}

	// Return the extracted URL in the JSON response
	ctx.JSON(http.StatusOK, gin.H{"Message": "Got the url as " + url})
}
