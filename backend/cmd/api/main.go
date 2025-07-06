package main

import (
	"log/slog"

	"github.com/doc4suresh/webPageAnalyzer/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error("No .env file found, using environment variables")
	}

	if err := server.Run(); err != nil {
		slog.Error("Failed to start the server", slog.Any("Error", err))
	}
}
