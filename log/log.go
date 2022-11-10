package log

import "context"

func Trace(args ...interface{}) {

}
func Tracef(args ...interface{}) {

}
func Debug(args ...interface{})                   {}
func Debugf(format string, args ...interface{})   {}
func Info(args ...interface{})                    {}
func Infof(format string, args ...interface{})    {}
func Warring(args ...interface{})                 {}
func Warringf(format string, args ...interface{}) {}
func Error(args ...interface{})                   {}
func Errorf(format string, args ...interface{})   {}

func WithContext(ctx context.Context) Logger {

}
