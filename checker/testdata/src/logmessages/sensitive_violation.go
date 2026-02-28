package logmessages

import (
	"go.uber.org/zap"
)

func ExampleZapSensitiveViolation() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	password := "123"
	apiKey := "abc"

	// want "log message may contain sensitive data"
	logger.Info("password: " + password)
	// want "log message may contain sensitive data"
	logger.Debug("api_key=" + apiKey)
}