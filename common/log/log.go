package log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// `gl` is the global logger.
	// Other packages should use public methods such as Info/Error to do the logging.
	// For other types of logging, e.g. logging to a separate file, they should use their own loggers.
	gl     *zap.Logger
	gLevel zap.AtomicLevel
)

// Initializes the global console logger.
func init() {
	gLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	gl, _ = zap.Config{
		Level:       gLevel,
		Development: true,
		// Use "console" to print readable stacktrace.
		Encoding: "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     customizedTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build(
		// Skip one caller stack to locate the correct caller.
		zap.AddCallerSkip(1),
	)
}

const (
	timeFormatLayout = "2006-01-02T15:04:05.000Z07:00"
)

// CustomizedTimeEncoder new a zap customized time encoder
func customizedTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(timeFormatLayout))
}

// SetLevel wraps the zap Level's SetLevel method.
func SetLevel(level zapcore.Level) {
	gLevel.SetLevel(level)
}

// EnabledLevel wraps the zap Level's Enabled method.
func EnabledLevel(level zapcore.Level) bool {
	return gLevel.Enabled(level)
}

// Debug wraps the zap Logger's Debug method.
func Debug(msg string, fields ...zap.Field) {
	gl.Debug(msg, fields...)
}

// Info wraps the zap Logger's Info method.
func Info(msg string, fields ...zap.Field) {
	gl.Info(msg, fields...)
}

// Warn wraps the zap Logger's Warn method.
func Warn(msg string, fields ...zap.Field) {
	gl.Warn(msg, fields...)
}

// Error wraps the zap Logger's Error method.
func Error(msg string, fields ...zap.Field) {
	gl.Error(msg, fields...)
}

// Sync wraps the zap Logger's Sync method.
func Sync() {
	_ = gl.Sync()
}
