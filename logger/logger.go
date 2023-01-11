package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)



type zapLoggerStruct struct {
	wrappedZapLogger *zap.Logger
}

func ZapLogger(debug bool) *zapLoggerStruct {
	logger := zap.NewDevelopmentConfig()

	if debug {
		logger.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		logger.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}
	logger.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger.DisableStacktrace = true
	logger.DisableCaller = true
	prettyLogger, _ := logger.Build()


	return &zapLoggerStruct{
		wrappedZapLogger: prettyLogger,
	}
}

func (z *zapLoggerStruct) MakeLoggerGlobal() {
	_ = zap.ReplaceGlobals(z.wrappedZapLogger)
}
