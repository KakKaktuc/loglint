package logmessages

import (
	"log/slog"
	"go.uber.org/zap"
)

func ExampleValidLogs() {
	slogLogger := slog.New(nil)
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()

	// корректные slog сообщения
	slogLogger.Info("starting server on port 8080")
	slogLogger.Error("failed to connect to database")

	// корректные zap сообщения
	zapLogger.Info("user authenticated successfully")
	zapLogger.Debug("api request completed")
}