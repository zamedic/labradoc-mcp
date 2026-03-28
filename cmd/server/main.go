package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/spf13/viper"

	"github.com/zamedic/labradoc-mcp/internal/labradoc"
	"github.com/zamedic/labradoc-mcp/internal/server"
)

func main() {
	// Configure viper for env var support
	viper.SetEnvPrefix("LABRADOC")
	viper.AutomaticEnv()

	viper.SetDefault("api_url", "https://labradoc.eu")
	viper.SetDefault("log_level", "info")
	viper.SetDefault("log_format", "text")

	// Load config
	apiKey := viper.GetString("api_key")
	apiURL := viper.GetString("api_url")
	logLevel := viper.GetString("log_level")
	logFormat := viper.GetString("log_format")

	// Configure logger
	var level slog.Level
	if err := level.UnmarshalText([]byte(logLevel)); err != nil {
		level = slog.LevelInfo
	}

	var logHandler slog.Handler
	if logFormat == "json" {
		logHandler = slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: level})
	} else {
		logHandler = slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level})
	}
	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	if apiKey == "" {
		logger.Error("LABRADOC_API_KEY is required")
		os.Exit(1)
	}

	logger.Info("Starting Labradoc MCP server",
		slog.String("api_url", apiURL),
		slog.String("log_level", logLevel),
	)

	// Create Labradoc client
	client := labradoc.NewClient(apiKey, apiURL, logger)

	// Create and run MCP server
	ctx := context.Background()
	mcpServer := server.NewMCPServer(client, logger)

	if err := mcpServer.Run(ctx); err != nil {
		logger.Error("Server exited with error", slog.Any("error", err))
		os.Exit(1)
	}
}
