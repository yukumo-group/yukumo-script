package logger_test

import (
	"testing"

	"github.com/yukumo-group/yukumo-script/pkg/utils/logger"
)

// Test the logging
func TestLogging(t *testing.T) {
	logger1 := logger.NewLogger(
		"test",
		&t,
	)
	logger1.Info("Test!")
	logger1.Error("Test!")
	loggerNil := logger.NewLogger(
		"test",
		nil,
	)
	loggerNil.Info("Test!")
	loggerNil.Error("Test!")
}
