package main

import (
	"log/slog"
	"os"
)

func main() {
	// Set the log level to debug if the DEBUG environment variable is set to true
	if is_debug := os.Getenv("DEBUG"); is_debug == "true" {
		slog.Warn("DEBUG MODE ENABLED")
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	r := startRouter()

	// TODO: Add this to a configuration file
	r.Run(":8080")
}
