package logger

import (
	"go.uber.org/zap"
)

// Logger is wrapper struct around logger.Logger that adds some custom functionality
type Logger struct {
	*zap.SugaredLogger
}

// NewLogger return logger instance
func NewLogger() *Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	return &Logger{sugar}
}
