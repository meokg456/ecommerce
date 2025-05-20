package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewAppLogger() (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()

	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	config.EncoderConfig.CallerKey = "func"
	config.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}
