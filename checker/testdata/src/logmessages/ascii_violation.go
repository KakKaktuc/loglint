package logmessages

import "log/slog"

func ExampleSlogASCIIViolation() {
	logger := slog.New(nil)
	// want "log message must contain only english ascii characters with lowercase letters and allowed symbols"
	logger.Info("Привет мир") // кириллица вызывает ASCII ошибку
}