package logger

import (
	"context"
	glog "log"
	"sync"
	"time"
	_ "time"

	"github.com/xavicci/gRPC-ModernAPI/books-app/internal/pkg/request"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Log   *zap.Logger
	Sugar *zap.SugaredLogger

	customTimeFormat string

	onceInit sync.Once
)

func init() {
	if err := Init(); err != nil {
		glog.Fatalf("Failed to initialize logger: %v", err)
	}
}

func Init() error {
	var err error

	onceInit.Do(func() {

		cfg := zap.Config{
			Encoding:         "json",
			Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stdout"},
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:  "msg",
				LevelKey:    "level",
				EncodeLevel: zapcore.CapitalLevelEncoder,

				TimeKey: "timestamp",
				EncodeTime: zapcore.TimeEncoder(func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
					enc.AppendString(t.UTC().Format("2006-01-02T15:04:05Z0700"))
				}),

				CallerKey:    "caller",
				EncodeCaller: zapcore.ShortCallerEncoder,

				StacktraceKey: "stacktrace",
			},
		}
		Log, _ = cfg.Build()

		Sugar = Log.Sugar()
		zap.RedirectStdLog(Log)
	})
	return err
}

func addReqField(ctx context.Context) zapcore.Field {
	return zap.String(request.RequestIDKey, request.GetContextRequestID(ctx))
}

func WithContext(ctx context.Context) *zap.SugaredLogger {
	if ctx != nil {
		return Sugar.With(addReqField(ctx))
	}
	return Sugar
}
