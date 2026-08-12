package api

import (
	"github.com/yukumo-group/yukumo-script/pkg/utils/logger"
)

// NewLogger creates new logger through api
func NewLogger(
	serviceName string,
) *logger.Logger {
	return logger.NewLogger(
		serviceName,
		nil,
	)
}
