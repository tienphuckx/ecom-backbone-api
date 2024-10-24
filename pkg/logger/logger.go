package logger

import (
	"log"
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/tienphuckx/ecom-backbone-api.git/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LogConfig) *LoggerZap {
	logLevelOptions := config.LogLevel

	var logLevel zapcore.Level
	switch logLevelOptions {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "info":
		logLevel = zapcore.InfoLevel
	case "warn":
		logLevel = zapcore.WarnLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	case "fatal":
		logLevel = zapcore.FatalLevel
	case "panic":
		logLevel = zapcore.PanicLevel
	default:
		// Fallback to a default log level and log a warning
		logLevel = zapcore.InfoLevel
		log.Printf("Invalid log level option: %s, using default 'info' level", logLevelOptions)
	}

	// Rest of your logger initialization code
	encoder := getEncoderLog_Dev()
	hook := lumberjack.Logger{
		Filename:   config.FileLogName,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(&hook),
		),
		logLevel,
	)

	return &LoggerZap{
		Logger: zap.New(
			core,
			zap.AddCaller(),
			zap.AddStacktrace(zap.ErrorLevel),
		),
	}
}

func getEncoderLog_Dev() zapcore.Encoder {
	config := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// Create a JSON encoder
	encoder := zapcore.NewJSONEncoder(config)

	return encoder
}
