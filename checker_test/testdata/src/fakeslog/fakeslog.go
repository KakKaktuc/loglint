package fakeslog

type Logger struct{}

func (l *Logger) Info(msg string)  {}
func (l *Logger) Error(msg string) {}