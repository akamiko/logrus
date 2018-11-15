package logger

import (
	"context"
)

// Logger インタフェース定義
type Logger interface {
	Debug(ctx context.Context, format string, msg string)
	Info(ctx context.Context, format string, msg string)
	Warn(ctx context.Context, format string, msg string)
	Error(ctx context.Context, format string, msg string)
	Fatal(ctx context.Context, format string, msg string)
}
