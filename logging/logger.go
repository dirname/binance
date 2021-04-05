package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var sugaredLogger *zap.SugaredLogger
var atomicLevel zap.AtomicLevel

func init() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:     "time",
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime:  zapcore.RFC3339NanoTimeEncoder,
	}

	atomicLevel = zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapcore.DebugLevel)

	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), os.Stdout, atomicLevel)
	sugaredLogger = zap.New(core).Sugar()
}

// SetLevel set the logger level
func SetLevel(level zapcore.Level) {
	atomicLevel.SetLevel(level)
}

// Fatal output fatal level msg
func Fatal(format string, args ...interface{}) {
	sugaredLogger.Fatalf(format, args...)
}

// Error output error level msg
func Error(format string, args ...interface{}) {
	sugaredLogger.Errorf(format, args...)
}

// Panic output panic level msg
func Panic(format string, args ...interface{}) {
	sugaredLogger.Panicf(format, args...)
}

// DPanic output DPanic level msg
func DPanic(format string, args ...interface{}) {
	sugaredLogger.DPanicf(format, args...)
}

// Warn output waning level msg
func Warn(format string, args ...interface{}) {
	sugaredLogger.Warnf(format, args...)
}

// Info output info level msg
func Info(format string, args ...interface{}) {
	sugaredLogger.Infof(format, args...)
}

// Debug out debug level msg
func Debug(format string, args ...interface{}) {
	sugaredLogger.Debugf(format, args...)
}