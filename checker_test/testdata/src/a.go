package src

import (
	fakelog "github.com/KakKaktuc/loglint/checker_test/testdata/src/fakelog"
	fakeslog "github.com/KakKaktuc/loglint/checker_test/testdata/src/fakeslog"
)

func main() {
	logger := &fakelog.Logger{}
	slog := &fakeslog.Logger{}
	password := "123"
	apiKey := "abc"

	// ❌ неправильные логи
	slog.Info("Starting server")                // want "log message must contain only lowercase english letters and allowed characters"
	slog.Info("запуск сервера")                // want "log message must contain only english ascii characters with lowercase letters and allowed symbols"
	slog.Info("server started 🚀")             // want "log message must contain only english ascii characters with lowercase letters and allowed symbols"
	logger.Info("user password: " + password)  // want "log message may contain sensitive data"
	logger.Debug("api_key=" + apiKey)          // want "log message may contain sensitive data"

	// ✅ правильные логи
	slog.Info("starting server")
	slog.Error("failed to connect to database")
	logger.Info("token validated")
}