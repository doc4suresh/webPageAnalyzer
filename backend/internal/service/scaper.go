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
	if err := CheckConnection(url); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to reach the URL"})
		return
	}

	info, err := GetWebPageInfo(url)
	if err != nil {
		log.Printf("Error getting web page info for %s: %v", url, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get web page info"})
		return
	}

	ctx.JSON(http.StatusOK, info)
}

func CheckConnection(url string) error {
	// Create a new collector
	collector := colly.NewCollector()

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

func GetWebPageInfo(url string) (*WebPageInfo, error) {

	info := &WebPageInfo{
		URL:       url,
		HeadCount: make(map[string]int),
	}

	collector := colly.NewCollector()

	// Extract the HTML version from the HTML
	collector.OnHTML("html", func(e *colly.HTMLElement) {
		info.HtmlVersion = e.Attr("html")
	})

	// Extract the title from the HTML
	collector.OnHTML("title", func(e *colly.HTMLElement) {
		info.Title = e.Text
	})

	// Extract the head count for each heading level
	collector.OnHTML("h1, h2, h3, h4, h5, h6", func(e *colly.HTMLElement) {
		info.HeadCount[e.Name]++
	})

	// Extract the accessible links
	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		info.AccessbleLinks = append(info.AccessbleLinks, e.Attr("href"))
	})

	// Extract the inaccessble links
	collector.OnHTML("a[href^='javascript:']", func(e *colly.HTMLElement) {
		info.InaccessbleLinks = append(info.InaccessbleLinks, e.Attr("href"))
	})

	// Extract the login form
	collector.OnHTML("form[action^='login']", func(e *colly.HTMLElement) {
		info.IsLogingForm = true
	})

	if err := collector.Visit(url); err != nil {
		return nil, err
	}

	return info, nil
}
