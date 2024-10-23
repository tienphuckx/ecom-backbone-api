package main

import (
	"log"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	/*
		// Create a logger using the production settings, which outputs logs in JSON format.
		// This is generally used in production environments.
		logger, err := zap.NewProduction()
		if err != nil {
			panic(err) // Handle the error in case logger creation fails
		}
		defer logger.Sync() // Flushes buffer, if any, before the program exits.

		// Simple logging with various levels of severity
		logger.Debug("This is a debug message")  // Used for verbose messages that are helpful during debugging
		logger.Info("This is an info message")   // General informational messages
		logger.Warn("This is a warning")         // Warnings about potential issues
		logger.Error("This is an error message") // Log errors that prevent a certain operation but don't crash the program

		// Logging with fields to provide more context to the message.
		logger.Info("User logged in",
			zap.String("username", "john_doe"), // String field
			zap.Int("attempt", 1),              // Integer field
			zap.Bool("success", true),          // Boolean field
		)

		// Structured logging of time and duration
		logger.Info("Task completed",
			zap.Time("time", time.Now()),           // Logs the current time
			zap.Duration("elapsed", time.Second*2), // Logs a duration
		)

		// Log an error and automatically include a stack trace
		logger.Error("Failed to fetch data", zap.Error(err))

		// Example with a Sugared Logger for easier, more concise logging.
		// SugaredLogger is less performant but allows more flexible, formatted logging like `Printf`.
		sugar := logger.Sugar()
		sugar.Infof("This is an info log with formatting: %s has logged in", "john_doe")
		sugar.Warnf("Warning: Too many failed login attempts for user: %s", "john_doe")

		// Contextual logging: Create a child logger with preset fields for repeated use
		dbLogger := logger.With(zap.String("component", "database"))
		dbLogger.Info("Database connected successfully")
		dbLogger.Error("Failed to connect to database", zap.String("db_name", "users_db"))

		// Using custom log levels: atomic log level allows dynamic level changes
		// atomicLevel := zap.NewAtomicLevel()                                                  // Create atomic level for log level control
		// atomicLogger := zap.New(zap.NewProductionEncoderConfig().EncoderConfig, atomicLevel) // Custom logger with atomic level control
		// atomicLogger.Info("This log uses atomic level")

		// Sampling: Reduce log volume by logging only some of the messages based on rules
		// Zap allows configuring sampling to limit the number of logs emitted (e.g., only 1 out of N)
		sampledLogger := zap.NewExample() // Example logger that can simulate different setups
		sampledLogger.Info("Sampled logging setup example")

		// Adding caller information (file and line where the log originated)
		loggerWithCaller := logger.WithOptions(zap.AddCaller())
		loggerWithCaller.Info("This log includes caller info") // Logs with the file and line number

		// Logging stack traces when severe errors occur
		loggerWithCaller.Error("Critical failure", zap.Stack("stacktrace"))

		// Create log with a custom field (you can log arbitrary types using zap.Any)
		logger.Info("Logging custom fields", zap.Any("customObject", map[string]interface{}{
			"id":    101,
			"name":  "Sample Object",
			"value": 42,
		}))

		// Graceful shutdown: flush the logs before exit (done with logger.Sync())
		// In practice, you should ensure all logs are flushed, especially in buffered logging setups.
		logger.Sync()


	*/

	encoder := getEncoderLog_Dev()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Dev: Info log", zap.Int("line", 1))
	logger.Error("Dev: Error log", zap.Int("line", 2))

	encoder_pro := getEncoderLog_Production()
	sync_pro := getWriterSync()
	core_pro := zapcore.NewCore(encoder_pro, sync_pro, zapcore.InfoLevel)
	logger_pro := zap.New(core_pro, zap.AddCaller())

	logger_pro.Info("Production: Info log", zap.Int("line", 1))
	logger_pro.Error("Production: Error log", zap.Int("line", 2))

}

// CUSTOM
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

func getEncoderLog_Production() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	//12323123.434234 -> "2023-01-23T12:32:31.434234Z"
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Time key: "timestamp" -> "timestamp" in log output (Default is "ts")
	encodeConfig.TimeKey = "timestamp"

	//"info" -> "INFO"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// Line number
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	// Message key: "message" -> "message" in log output (Default is "msg")
	encodeConfig.MessageKey = "message"

	// Stack trace key: "stacktrace" -> "stacktrace" in log output (Default is "stacktrace")
	encodeConfig.StacktraceKey = "stacktrace"

	// Create a JSON encoder
	encoder := zapcore.NewJSONEncoder(encodeConfig)

	return encoder
}

func getWriterSync_error_create_log() zapcore.WriteSyncer {
	file, _ := os.OpenFile(".log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}

func getWriterSync() zapcore.WriteSyncer {
	// Ensure the directory exists
	logDir := ".log"
	logFilePath := filepath.Join(logDir, "log.txt")

	// Create the log directory if it doesn't exist
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	// Open the log file
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Create write syncers for both file and console (stderr)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)

	// Write logs to both file and console
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}

// Function to create the logger with different outputs for normal and error logs
func newLogger() *zap.Logger {
	// Encoder config for JSON format (you can also use console encoders if needed)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Use human-readable time format

	// Create encoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// Normal log file (info, debug, etc.)
	file, _ := os.OpenFile("./log/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	syncFile := zapcore.AddSync(file)

	// Error log file
	errorFile, _ := os.OpenFile("./log/error_log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	syncErrorFile := zapcore.AddSync(errorFile)

	// Console output
	syncConsole := zapcore.AddSync(os.Stderr)

	// Combine normal logs to both console and the general log file
	normalWriteSyncer := zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
	normalCore := zapcore.NewCore(encoder, normalWriteSyncer, zapcore.InfoLevel)

	// Combine error logs to both console and the error log file
	errorWriteSyncer := zapcore.NewMultiWriteSyncer(syncConsole, syncErrorFile)
	errorCore := zapcore.NewCore(encoder, errorWriteSyncer, zapcore.ErrorLevel)

	// Combine both cores
	core := zapcore.NewTee(normalCore, errorCore)

	// Return the logger
	return zap.New(core)
}
