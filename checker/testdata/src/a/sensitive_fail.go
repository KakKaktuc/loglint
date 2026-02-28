package a

import "log/slog"

func SensitiveFail(l *slog.Logger) {
    l.Info("token=xyz") // want "log message may contain sensitive data"
}