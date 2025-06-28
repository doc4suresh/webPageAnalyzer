package main

import (
	"log"

	"github.com/doc4suresh/webPageAnalyzer/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	if err := server.Run(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
