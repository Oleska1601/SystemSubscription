package controller

type LoggerTest struct{}

func (l *LoggerTest) Debug(msg string, args ...any) {}
func (l *LoggerTest) Error(msg string, args ...any) {}
func (l *LoggerTest) Info(msg string, args ...any)  {}
func (l *LoggerTest) Warn(msg string, args ...any)  {}
