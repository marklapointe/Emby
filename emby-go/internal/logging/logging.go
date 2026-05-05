package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger creates a configured zap logger based on the given level and format.
func NewLogger(level, format string) (*zap.Logger, error) {
	var cfg zapcore.EncoderConfig
	var levelCfg zapcore.Level

	if err := levelCfg.UnmarshalText([]byte(level)); err != nil {
		levelCfg = zapcore.InfoLevel
	}

	// Encoder config
	cfg.TimeKey       = "ts"
	cfg.LevelKey      = "level"
	cfg.NameKey       = "logger"
	cfg.CallerKey     = "caller"
	cfg.MessageKey    = "msg"
	cfg.StacktraceKey = "stacktrace"
	cfg.EncodeLevel   = zapcore.CapitalLevelEncoder
	cfg.EncodeTime    = zapcore.ISO8601TimeEncoder
	cfg.EncodeCaller  = zapcore.ShortCallerEncoder

	var encoder zapcore.Encoder
	if format == "console" {
		encoder = zapcore.NewConsoleEncoder(cfg)
	} else {
		encoder = zapcore.NewJSONEncoder(cfg)
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(os.Stdout),
		levelCfg,
	)

	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)), nil
}

// NewSilentLogger returns a logger that discards all output (useful for testing).
func NewSilentLogger() *zap.Logger {
	return zap.NewNop()
}
