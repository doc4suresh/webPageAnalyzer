package service

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func Scrape(ctx *gin.Context, url string) {
	// Create a new collector
	collector := colly.NewCollector()

	if err := CheckConnection(collector, url); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to reach the URL"})
		return
	}
}

func CheckConnection(collector *colly.Collector, url string) error {
	// Parse retry delay from environment variable
	retryDelayStr := os.Getenv("URL_RETRY_DELAY")
	retryDelay, err := strconv.Atoi(retryDelayStr)
	if err != nil || retryDelay < 1 {
		retryDelay = 1 // default to 1 second if not set or invalid
	}

	// Parse retry limit from environment variable
	retryLimitStr := os.Getenv("URL_RETRY_LIMIT")
	retryLimit, err := strconv.Atoi(retryLimitStr)
	if err != nil || retryLimit < 1 {
		retryLimit = 1 // default to 1 if not set or invalid
	}

	// Check if the URL is reachable
	for i := 0; i < retryLimit; i++ {
		err := collector.Visit(url)

		if err != nil {
			log.Printf("Attempt %d: Failed to reach the URL %s, error: %v", i+1, url, err)
			time.Sleep(time.Duration(retryDelay) * time.Second)
		} else {
			log.Printf("Successfully reached the URL %s on attempt %d", url, i+1)
			return nil
		}
	}

	log.Printf("Failed to reach the URL %s after %d attempts", url, retryLimit)
	return fmt.Errorf("failed to reach the URL %s after %d attempts", url, retryLimit)
}
