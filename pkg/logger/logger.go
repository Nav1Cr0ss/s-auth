package logger

import (
	"github.com/Nav1Cr0ss/s-auth/internal/config"
	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
}

// NewLogger return logger instance
func NewLogger(cfg *config.Config) *Logger {
	loggers := map[bool]func(options ...zap.Option) (*zap.Logger, error){
		true:  zap.NewDevelopment,
		false: zap.NewProduction,
	}

	logger, _ := loggers[cfg.Debug]()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	return &Logger{sugar}
}
