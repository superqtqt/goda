package log

import (
	"context"
	"sync"
)

type Level int

const (
	LevelNil Level = iota
	LevelTrace
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
)

var (
	DefaultLogger     Logger
	mu                sync.RWMutex
	loggers           = make(map[string]Logger)
	DefaultLoggerName = "default"
)

func init() {
	//TODO init default logger
}

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

func Register(name string, logger Logger) {
	mu.Lock()
	defer mu.Unlock()
	if logger == nil {
		panic("log: Register logger is nil")
	}
	if _, dup := loggers[name]; dup && name != DefaultLoggerName {
		panic("log: Register called twiced for logger name " + name)
	}
	loggers[name] = logger
	if name == DefaultLoggerName {
		DefaultLogger = logger
	}
}

// GetDefaultLogger gets the default Logger.
func GetDefaultLogger() Logger {
	mu.RLock()
	l := DefaultLogger
	mu.RUnlock()
	return l
}

func SetDefaultLogger(logger Logger) {
	mu.Lock()
	DefaultLogger = logger
	mu.Unlock()
}

// Get returns the Logger implementation by log name.
func Get(name string) Logger {
	mu.RLock()
	l := loggers[name]
	mu.RUnlock()
	return l
}
