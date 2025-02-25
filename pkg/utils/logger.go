package utils

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	clt *zap.Logger
}

func NewLogger() Logger {
	cfg := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig = encoderConfig

	l, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	return Logger{clt: l}
}

func (logger *Logger) Info(message any, fields ...zap.Field) {
	logger.clt.Info(fmt.Sprintf("%#v", message), fields...)
}

func (logger *Logger) Debug(message string, fields ...zap.Field) {
	logger.clt.Debug(message, fields...)
}

func (logger *Logger) Error(message string, fields ...zap.Field) {
	logger.clt.Error(message, fields...)
}
