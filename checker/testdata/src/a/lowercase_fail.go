package a

import "log/slog"

func LowercaseFail(l *slog.Logger) {
    l.Info("Bad Message!") // want "log message must contain only lowercase english letters"
}