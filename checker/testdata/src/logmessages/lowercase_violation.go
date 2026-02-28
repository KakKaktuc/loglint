package logmessages

import "log/slog"

func ExampleSlogLowercaseViolation() {
	logger := slog.New(nil)
	// want "log message must contain only lowercase english letters and allowed characters"
	logger.Info("Hello World!") // заглавные буквы и ! вызывают правило lowercase
}