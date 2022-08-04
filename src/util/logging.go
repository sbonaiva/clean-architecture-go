package util

import "go.uber.org/zap"

type logger struct {
	tool *zap.SugaredLogger
}

type Logger interface {
	Debug(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
	Fatal(msg string, keysAndValues ...interface{})
	Panic(msg string, keysAndValues ...interface{})
}

func NewLogger(t *zap.SugaredLogger) Logger {
	return &logger{
		tool: t,
	}
}

func (l *logger) Debug(msg string, keysAndValues ...interface{}) {
	l.tool.Debugw(msg, keysAndValues...)
}

func (l *logger) Info(msg string, keysAndValues ...interface{}) {
	l.tool.Infow(msg, keysAndValues...)
}

func (l *logger) Warn(msg string, keysAndValues ...interface{}) {
	l.tool.Warnw(msg, keysAndValues...)
}

func (l *logger) Error(msg string, keysAndValues ...interface{}) {
	l.tool.Errorw(msg, keysAndValues...)
}

func (l *logger) Fatal(msg string, keysAndValues ...interface{}) {
	l.tool.Fatalw(msg, keysAndValues...)
}

func (l *logger) Panic(msg string, keysAndValues ...interface{}) {
	l.tool.Panicw(msg, keysAndValues...)
}
