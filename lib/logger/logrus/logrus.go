package logrus

import (
	"context"

	"github.com/akamiko/logging-sample4/lib/logger"
	"github.com/sirupsen/logrus"
)

// Logger 構造体定義
type Logger struct {
	log *logrus.Logger
}

// NewLogger インスタンス生成
func NewLogger() logger.Logger {
	log := logrus.New()
	log.SetFormatter(new(logrus.JSONFormatter))
	log.SetLevel(logrus.InfoLevel)
	return &Logger{log}
}

// Debug Debugレベルのログを出力する関数
func (logger *Logger) Debug(ctx context.Context, format string, msg string) {
	logger.log.Debugf(format, msg)
}

// Info Infoレベルのログを出力する関数
func (logger *Logger) Info(ctx context.Context, format string, msg string) {
	logger.log.Infof(format, msg)
}

// Warn Warnレベルのログを出力する関数
func (logger *Logger) Warn(ctx context.Context, format string, msg string) {
	logger.log.Warnf(format, msg)
}

// Error Errorレベルのログを出力する関数
func (logger *Logger) Error(ctx context.Context, format string, msg string) {
	logger.log.Errorf(format, msg)
}

// Fatal Fatalレベルのログを出力する関数
func (logger *Logger) Fatal(ctx context.Context, format string, msg string) {
	logger.log.Fatalf(format, msg)
}
