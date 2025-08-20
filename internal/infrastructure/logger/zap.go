package logger

import (
	"clean-architecture/internal/domain/service/logger"
	"fmt"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type OutputType string

const (
	OutputTypeConsole OutputType = "console"
	OutputTypeFile    OutputType = "file"
)

type Option func(*options)

type options struct {
	level        zapcore.Level
	outputs      []OutputType
	logFilePath  string
	timeFormat   string
	customLevels bool
}

type Client interface {
	Debug(msg string, fields ...logger.LogField)
	Info(msg string, fields ...logger.LogField)
	Warn(msg string, fields ...logger.LogField)
	Error(msg string, fields ...logger.LogField)
	Fatal(msg string, fields ...logger.LogField)
}

type client struct {
	logger *zap.Logger
}

func buildOptions(opts ...Option) *options {
	dopts := &options{
		level:      zapcore.InfoLevel,
		outputs:    []OutputType{OutputTypeConsole},
		timeFormat: time.RFC3339,
	}
	for i := range opts {
		opts[i](dopts)
	}
	return dopts
}

func WithLevel(level string) Option {
	return func(opt *options) {
		if err := opt.level.UnmarshalText([]byte(level)); err != nil {
			opt.level = zapcore.InfoLevel
		}
	}
}

func WithConsole() Option {
	return func(opt *options) {
		for _, out := range opt.outputs {
			if out == OutputTypeConsole {
				return
			}
		}
		opt.outputs = append(opt.outputs, OutputTypeConsole)
	}
}

func WithFile(path string) Option {
	return func(opt *options) {
		opt.outputs = append(opt.outputs, OutputTypeFile)
		opt.logFilePath = path
	}
}

func WithTimeFormat(format string) Option {
	return func(opt *options) {
		opt.timeFormat = format
	}
}

func WithCustomLevels() Option {
	return func(opt *options) {
		opt.customLevels = true
	}
}

// NewClient
func NewClient(opts ...Option) (Client, error) {
	options := buildOptions(opts...)

	// Encoder config
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    encodeLevel,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(options.timeFormat))
		},
	}

	// Build cores
	var cores []zapcore.Core
	for _, out := range options.outputs {
		switch out {
		case OutputTypeConsole:
			if !options.customLevels {
				encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
			}
			encoder := zapcore.NewConsoleEncoder(encoderCfg)
			cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), options.level))
		case OutputTypeFile:
			if options.logFilePath == "" {
				return nil, fmt.Errorf("log file path is empty")
			}
			file, err := os.OpenFile(options.logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
			if err != nil {
				return nil, fmt.Errorf("failed to open log file: %w", err)
			}
			fileEncoder := zapcore.NewJSONEncoder(encoderCfg)
			cores = append(cores, zapcore.NewCore(fileEncoder, zapcore.AddSync(file), options.level))
		}
	}

	// Merge outputs
	core := zapcore.NewTee(cores...)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return &client{logger: logger}, nil
}

func encodeLevel(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var severity string
	switch l {
	case zapcore.DebugLevel:
		severity = "DEBUG"
	case zapcore.InfoLevel:
		severity = "INFO"
	case zapcore.WarnLevel:
		severity = "WARNING"
	case zapcore.ErrorLevel:
		severity = "ERROR"
	case zapcore.DPanicLevel:
		severity = "CRITICAL"
	case zapcore.PanicLevel:
		severity = "ALERT"
	case zapcore.FatalLevel:
		severity = "EMERGENCY"
	default:
		severity = "INFO"
	}
	enc.AppendString(strings.ToUpper(severity))
}

func (c *client) Debug(msg string, fields ...logger.LogField) {
	c.logger.Debug(msg, convertFields(fields)...)
}

func (c *client) Info(msg string, fields ...logger.LogField) {
	c.logger.Info(msg, convertFields(fields)...)
}

func (c *client) Warn(msg string, fields ...logger.LogField) {
	c.logger.Warn(msg, convertFields(fields)...)
}

func (c *client) Error(msg string, fields ...logger.LogField) {
	c.logger.Error(msg, convertFields(fields)...)
}

func (c *client) Fatal(msg string, fields ...logger.LogField) {
	c.logger.Fatal(msg, convertFields(fields)...)
}

func convertFields(fields []logger.LogField) []zap.Field {
	zFields := make([]zap.Field, len(fields))
	for i, field := range fields {
		switch field.Type {
		case logger.LogTypeString:
			zFields[i] = zap.String(field.Key, field.String)
		case logger.LogTypeInt:
			zFields[i] = zap.Int(field.Key, int(field.Integer))
		case logger.LogTypeError:
			zFields[i] = zap.Error(field.Err)
		case logger.LogTypeAny:
			zFields[i] = zap.Any(field.Key, field.Interface)
		default:
			zFields[i] = zap.Any(field.Key, field.Interface)
		}
	}
	return zFields
}
