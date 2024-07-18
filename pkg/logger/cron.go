package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type CronLogger struct {
	zap *zap.Logger
}

func NewCronLogger() *CronLogger {
	return &CronLogger{
		zap: Zap(),
	}
}

func (l *CronLogger) Info(msg string, keysAndValues ...interface{}) {
	var fields []zapcore.Field

	if len(keysAndValues)%2 == 0 {
		for i := 0; i < len(keysAndValues); i += 2 {
			fields = append(fields, zap.Any(keysAndValues[i].(string), keysAndValues[i+1]))
		}
	} else {
		fields = append(fields, zap.Any("msg", keysAndValues))
	}

	l.zap.Info(msg, fields...)
}

func (l *CronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	var fields []zapcore.Field

	fields = append(fields, zap.Error(err))

	if len(keysAndValues)%2 == 0 {
		for i := 0; i < len(keysAndValues); i += 2 {
			fields = append(fields, zap.Any(keysAndValues[i].(string), keysAndValues[i+1]))
		}
	} else {
		fields = append(fields, zap.Any("msg", keysAndValues))
	}

	l.zap.Error(msg, fields...)
}
