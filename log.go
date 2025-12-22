package iamlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// zapLogger is a logr.Logger that uses Zap to log.
type zapLogger struct {
	// NB: this looks very similar to zap.SugaredLogger, but
	// deals with our desire to have multiple verbosity levels.
	zapLogger *zap.Logger
}

var (
	logger  *zapLogger
	options *Options
)

// nolint: gochecknoinits // need to init a default logger
func init() {
	Init(NewOptions())
}

// Init initializes logger by opts which can custmoized by command arguments.
func Init(opts *Options) {
	options = opts
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: milliSecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	// when output to local path, with color is forbidden
	if opts.EnableColor {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = InfoLevel
	}

	loggerConfig := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       false,
		DisableCaller:     !opts.EnableCaller,
		DisableStacktrace: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         opts.Format,
		EncoderConfig:    encoderConfig,
		OutputPaths:      opts.OutputPaths,
		ErrorOutputPaths: opts.ErrorOutputPaths,
	}

	var err error
	l, err := loggerConfig.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	logger = &zapLogger{
		zapLogger: l,
	}
	zap.RedirectStdLog(l)
}

// Debug method output debug level log.
func Debug(msg string, fields ...Field) {
	logger.zapLogger.Debug(msg, fields...)
}

// Debugf method output debug level log.
func Debugf(format string, v ...interface{}) {
	logger.zapLogger.Sugar().Debugf(format, v...)
}

// Debugw method output debug level log.
func Debugw(msg string, keysAndValues ...interface{}) {
	logger.zapLogger.Sugar().Debugw(msg, keysAndValues...)
}

// Info method output info level log.
func Info(msg string, fields ...Field) {
	logger.zapLogger.Info(msg, fields...)
}

// Infof method output info level log.
func Infof(format string, v ...interface{}) {
	logger.zapLogger.Sugar().Infof(format, v...)
}

// Infow method output info level log.
func Infow(msg string, keysAndValues ...interface{}) {
	logger.zapLogger.Sugar().Infow(msg, keysAndValues...)
}

// Warn method output warning level log.
func Warn(msg string, fields ...Field) {
	logger.zapLogger.Warn(msg, fields...)
}

// Warnf method output warning level log.
func Warnf(format string, v ...interface{}) {
	logger.zapLogger.Sugar().Warnf(format, v...)
}

// Warnw method output warning level log.
func Warnw(msg string, keysAndValues ...interface{}) {
	logger.zapLogger.Sugar().Warnw(msg, keysAndValues...)
}

// Error method output error level log.
func Error(msg string, fields ...Field) {
	logger.zapLogger.Error(msg, fields...)
}

// Errorf method output error level log.
func Errorf(format string, v ...interface{}) {
	logger.zapLogger.Sugar().Errorf(format, v...)
}

// Errorw method output error level log.
func Errorw(msg string, keysAndValues ...interface{}) {
	logger.zapLogger.Sugar().Errorw(msg, keysAndValues...)
}

// Panic method output panic level log and shutdown application.
func Panic(msg string, fields ...Field) {
	logger.zapLogger.Panic(msg, fields...)
}

// Panicf method output panic level log and shutdown application.
func Panicf(format string, v ...interface{}) {
	logger.zapLogger.Sugar().Panicf(format, v...)
}

// Panicw method output panic level log.
func Panicw(msg string, keysAndValues ...interface{}) {
	logger.zapLogger.Sugar().Panicw(msg, keysAndValues...)
}

// Fatal method output fatal level log.
func Fatal(msg string, fields ...Field) {
	logger.zapLogger.Fatal(msg, fields...)
}

// Fatalf method output fatal level log.
func Fatalf(format string, v ...interface{}) {
	logger.zapLogger.Sugar().Fatalf(format, v...)
}

// Fatalw method output Fatalw level log.
func Fatalw(msg string, keysAndValues ...interface{}) {
	logger.zapLogger.Sugar().Fatalw(msg, keysAndValues...)
}

func GetOptions() *Options {
	return options
}

func GetLogger() *zapLogger {
	return logger
}
