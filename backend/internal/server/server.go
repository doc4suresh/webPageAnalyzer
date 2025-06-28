package server

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Run initializes the Gin router, sets up middleware and routes, and starts the HTTP server.
// It returns an error if the server fails to start.
func Run() error {
	router := gin.New()
	router.Use(cors.Default())

	router.GET("/analyze", analyzeHandler)

	if err := router.Run(os.Getenv("SERVER_ADDRESS") + ":" + os.Getenv("SERVER_PORT")); err != nil {
		return fmt.Errorf("cound't start the backend server: %v", err)
	}

	return nil
}
