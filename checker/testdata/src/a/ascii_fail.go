package a

import "log/slog"

func AsciiFail(l *slog.Logger) {
    l.Info("привет мир") // want "log message must contain only english ascii characters"
}