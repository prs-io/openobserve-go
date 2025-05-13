package client

import (
	"log"
	"os"
)

// Logger is a simple logger interface that can be implemented for different logging strategies.
type Logger interface {
	Info(msg string)
	Error(msg string)
	Debug(msg string)
}

// DefaultLogger is a basic implementation of the Logger interface that logs to standard output.
type DefaultLogger struct {
	debug bool
}

// NewDefaultLogger creates a new instance of DefaultLogger.
func NewDefaultLogger(debug bool) *DefaultLogger {
	return &DefaultLogger{debug: debug}
}

// Info logs an informational message.
func (l *DefaultLogger) Info(msg string) {
	log.Println("INFO:", msg)
}

// Error logs an error message.
func (l *DefaultLogger) Error(msg string) {
	log.Println("ERROR:", msg)
}

// Debug logs a debug message if debugging is enabled.
func (l *DefaultLogger) Debug(msg string) {
	if l.debug {
		log.Println("DEBUG:", msg)
	}
}

// SetLogger allows the user to set a custom logger.
var SetLogger func(Logger)

// InitLogger initializes the logger with a default logger or a custom one.
func InitLogger(debug bool) {
	if SetLogger == nil {
		SetLogger = func(l Logger) {
			log.SetOutput(os.Stdout)
		}
	}
	SetLogger(NewDefaultLogger(debug))
}