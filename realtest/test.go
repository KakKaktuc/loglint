package main

import (
	"fmt"
	"log/slog"
	"os"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to sync logger: %v\n", err)
		}
	}()

	password := "test"
	apiKey := "test2"
	token := "test3"

	// ❌ неправильно
	slog.Info("Starting server on port 8080")
	slog.Error("Failed to connect to database")

	// ✅ правильно
	slog.Info("starting server on port 8080")
	slog.Error("failed to connect to database")

	// ❌ неправильно (русский)
	slog.Info("запуск сервера")

	// ❌ неправильно (emoji)
	slog.Info("server started 🚀")

	// ❌ неправильно (sensitive)
	logger.Info("user password: " + password)
	logger.Debug("api_key=" + apiKey)
	logger.Info("token: " + token)

	// ✅ правильно
	logger.Info("user authenticated successfully")
	logger.Debug("api request completed")
	logger.Info("token validated")
}
