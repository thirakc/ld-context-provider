package logz

import "go.uber.org/zap"

var logger = zap.NewExample()

func NewLogger() *zap.Logger {
	return logger
}
