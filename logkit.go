package logkit

import "log"

// Logger is a standard interface for logging messages.
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

// DefaultLogger implements Logger interface with standard log package.
type DefaultLogger struct{}

// Debug prints debug message.
func (dl *DefaultLogger) Debug(args ...interface{}) {
	log.Println(args...)
}

// Debugf prints formatted debug message.
func (dl *DefaultLogger) Debugf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// Info prints info message.
func (dl *DefaultLogger) Info(args ...interface{}) {
	log.Println(args...)
}

// Infof prints formatted info message.
func (dl *DefaultLogger) Infof(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// Warn prints warning message.
func (dl *DefaultLogger) Warn(args ...interface{}) {
	log.Println(args...)
}

// Warnf prints formatted warning message.
func (dl *DefaultLogger) Warnf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// Error prints error message.
func (dl *DefaultLogger) Error(args ...interface{}) {
	log.Println(args...)
}

// Errorf prints formatted error message.
func (dl *DefaultLogger) Errorf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// Fatal prints fatal message and exists with code 1.
func (dl *DefaultLogger) Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Fatalf prints formatted fatal message and exists with code 1.
func (dl *DefaultLogger) Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
