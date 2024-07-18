package logger

import (
	"i3/config"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func InitLogger() {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	consoleErrors := zapcore.Lock(os.Stderr)
	consoleDebugging := zapcore.Lock(os.Stdout)

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	cores := []zapcore.Core{
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
	}

	if config.ReadConfig().Env == "development" {
		cores = append(cores, zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority))
	}

	core := zapcore.NewTee(cores...)

	log = zap.New(core)
}

func Zap() *zap.Logger {
	if log == nil {
		InitLogger()
	}

	return log
}
