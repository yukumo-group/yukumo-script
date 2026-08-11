package logger

import (
	"fmt"
	"testing"
)

// Logger datastructure create logs for each service
type Logger struct {
	loggerName string
	testLogger **testing.T
}

// NewLogger creates a logger
func NewLogger(
	loggerName string,
	testLogger **testing.T, // testLogger have to pass in testing.T
) *Logger {
	return &Logger{
		loggerName: loggerName,
		testLogger: testLogger,
	}
}

// Info logs a info message
func (logger *Logger) Info(
	message string,
) {
	if syslogger != nil {
		syslogger.Info(
			fmt.Sprintf(
				"%s: %s",
				logger.loggerName,
				message,
			),
		)
	} else if logger.testLogger != nil {
		(*logger.testLogger).Logf(
			"[INFO]%s: %s",
			logger.loggerName,
			message,
		)
	}
}

// Error logs an error message
func (logger *Logger) Error(
	message string,
) {
	if syslogger != nil {
		syslogger.Error(
			fmt.Sprintf(
				"%s: %s",
				logger.loggerName,
				message,
			),
		)
	} else if logger.testLogger != nil {
		(*logger.testLogger).Logf(
			"[ERROR]%s: %s",
			logger.loggerName,
			message,
		)
	}
}
