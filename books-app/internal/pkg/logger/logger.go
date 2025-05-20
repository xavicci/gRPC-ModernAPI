package logger

import (
	_ "context"
	_ "sync"
	_ "time"

	_ "github.com/xavicci/gRPC-ModernAPI/books-app/internal/pkg/request"

	"go.uber.org/zap"
	_ "go.uber.org/zap/zapcore"
)

var (
	Log   *zap.Logger
	Sugar *zap.SugaredLogger
)
