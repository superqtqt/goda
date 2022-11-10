package log

import "context"

type Level int

const (
	LevelNil Level = iota
	LevelTrace
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
)

type Logger interface {
	Trace(args ...interface{})
	Tracef(args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warring(args ...interface{})
	Warringf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	WithContext(ctx context.Context) Logger
}

type LogConfig struct {
	Level Level
}
