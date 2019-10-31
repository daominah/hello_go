// Package log is a convenient way to call zap log funcs.

package zlog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var globalLogger *zap.SugaredLogger = newLogger()

func newLogger() *zap.SugaredLogger {
	var config zap.Config
	if os.Getenv("LOG_LEVEL_INFO") != "true" {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}
	config.Encoding = "console"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logFilePath := os.Getenv("LOG_FILE_PATH")
	if logFilePath != "" {
		fmt.Printf("app log to %v\n", logFilePath)
		config.OutputPaths = []string{logFilePath}
		config.ErrorOutputPaths = []string{logFilePath}
		if os.Getenv("LOG_NOT_STDERR") != "true" {
			config.OutputPaths = append(config.OutputPaths, "stderr")
			config.ErrorOutputPaths = append(config.ErrorOutputPaths, "stderr")
		}
	}
	// TODO: LOG_TIME_ROTATE

	wrapOption := zap.AddCallerSkip(1)
	zl, err := config.Build(wrapOption)
	if err != nil {
		log.Fatal("cannot build zap's logger", err)
	}
	defer zl.Sync()
	logger := zl.Sugar()
	return logger
}

func Fatal(args ...interface{}) {
	globalLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	globalLogger.Fatalf(template, args...)
}

func Info(args ...interface{}) {
	globalLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	globalLogger.Infof(template, args...)
}

func Debug(args ...interface{}) {
	globalLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	globalLogger.Debugf(template, args...)
}

func Println(args ...interface{}) {
	globalLogger.Info(args...)
}

func Printf(template string, args ...interface{}) {
	globalLogger.Infof(template, args...)
}
