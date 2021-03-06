package logger

import (
	"log"

	"go.uber.org/zap"
)

var zapLogger *zap.Logger

func init() {
	var err error
	zapLogger, err = zap.NewProduction()
	if err != nil {
		log.Fatal(err.Error())
	}
	Info("init logger ok")
}

func Info(msg string, fields ...zap.Field) {
	zapLogger.Info(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	zapLogger.Fatal(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	zapLogger.Error(msg, fields...)
}

func Sync() error {
	return zapLogger.Sync()
}
