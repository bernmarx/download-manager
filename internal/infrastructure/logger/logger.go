package logger

import "go.uber.org/zap"

var logger *zap.Logger

// Logger() provides access to a singleton logger
func Logger() *zap.Logger {
	if logger == nil {
		logger, _ = zap.NewProduction()
	}

	return logger
}
