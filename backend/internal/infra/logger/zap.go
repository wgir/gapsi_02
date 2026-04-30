package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

func NewLogger(env string) (*Logger, error) {
	var config zap.Config
	if env == "prod" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	l, err := config.Build()
	if err != nil {
		return nil, err
	}

	return &Logger{l}, nil
}

type contextKey string

const requestIDKey contextKey = "request_id"

func (l *Logger) WithRequestID(ctx context.Context) *zap.Logger {
	if requestID, ok := ctx.Value(requestIDKey).(string); ok {
		return l.With(zap.String("request_id", requestID))
	}
	return l.Logger
}

func ContextWithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}
