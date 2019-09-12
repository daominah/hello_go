package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func NewLogger() *zap.SugaredLogger {
	var zl *zap.Logger
	conf := zap.NewProductionConfig()
	conf.Encoding = "console"
	conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zl, _ = conf.Build()
	defer zl.Sync()
	logger := zl.Sugar()
	return logger
}

func Test1(t *testing.T) {
	log := NewLogger()
	log.Infof("ahihi: %v", "l1\nl2")
}
