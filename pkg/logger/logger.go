package logger

import (
	"os"
	"time"

	"auth/pkg/config"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(cfg *config.AuthConfig) logr.Logger {
	logLevel := zaplogLevel(cfg)
	logWriter := zapcore.AddSync(os.Stderr)

	var encoderCfg zapcore.EncoderConfig
	if cfg.Mode == config.Development {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	var encoder zapcore.Encoder
	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"
	encoderCfg.EncodeTime = syslogTimeEncoder

	if cfg.LogEncoding == "" && cfg.Mode == config.Development {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	newLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	// nolint // ignoring err
	newLogger.Sync()

	return zapr.NewLogger(newLogger)
}

func zaplogLevel(cfg *config.AuthConfig) (logLevel zapcore.Level) {
	switch cfg.LogLevel {
	case "debug":
		logLevel = zapcore.DebugLevel

	case "info":
		logLevel = zapcore.InfoLevel

	case "warn":
		logLevel = zapcore.WarnLevel

	case "error":
		logLevel = zapcore.ErrorLevel

	default:
		logLevel = zapcore.DebugLevel
	}

	return logLevel
}

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05"))
}
