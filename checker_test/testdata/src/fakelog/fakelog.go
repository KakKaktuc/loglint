package fakelog

type Logger struct{}

func (l *Logger) Info(msg string) {}
func (l *Logger) Error(msg string) {}
func (l *Logger) Debug(msg string) {}