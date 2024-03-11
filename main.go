package main

import (
	"Journal-Connector/logging"
	"Journal-Connector/parsing"
	"context"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
)

var log = logging.Logger{Package: "main"}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Load env vars
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Start the processing
	parsing.SetDebug(parsing.DebugLevelWarn)
	go parsing.StartProcessingAllFiles()

	<-ctx.Done()

	log.Println("Shutting down")
}
