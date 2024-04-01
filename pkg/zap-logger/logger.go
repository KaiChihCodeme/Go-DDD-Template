package zaplogger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// global vars
var logger *zap.Logger
var file *os.File
var logRotation *lumberjack.Logger

// global types
type Level = zapcore.Level
type Logger = zap.Logger
type Field = zap.Field
type WriteSyncer = zapcore.WriteSyncer

// const use for caller
const (
	InfoLevel  = zap.InfoLevel
	WarnLevel  = zap.WarnLevel
	ErrorLevel = zap.ErrorLevel
	DebugLevel = zap.DebugLevel
	PanicLevel = zap.PanicLevel
)

// fields type for caller
var (
	Skip        = zap.Skip
	Binary      = zap.Binary
	Bool        = zap.Bool
	Boolp       = zap.Boolp
	ByteString  = zap.ByteString
	Complex128  = zap.Complex128
	Complex128p = zap.Complex128p
	Complex64   = zap.Complex64
	Complex64p  = zap.Complex64p
	Float64     = zap.Float64
	Float64p    = zap.Float64p
	Float32     = zap.Float32
	Float32p    = zap.Float32p
	Int         = zap.Int
	Intp        = zap.Intp
	Int64       = zap.Int64
	Int64p      = zap.Int64p
	Int32       = zap.Int32
	Int32p      = zap.Int32p
	Int16       = zap.Int16
	Int16p      = zap.Int16p
	Int8        = zap.Int8
	Int8p       = zap.Int8p
	String      = zap.String
	Stringp     = zap.Stringp
	Uint        = zap.Uint
	Uintp       = zap.Uintp
	Uint64      = zap.Uint64
	Uint64p     = zap.Uint64p
	Uint32      = zap.Uint32
	Uint32p     = zap.Uint32p
	Uint16      = zap.Uint16
	Uint16p     = zap.Uint16p
	Uint8       = zap.Uint8
	Uint8p      = zap.Uint8p
	Uintptr     = zap.Uintptr
	Uintptrp    = zap.Uintptrp
	Reflect     = zap.Reflect
	Namespace   = zap.Namespace
	Stringer    = zap.Stringer
	Time        = zap.Time
	Timep       = zap.Timep
	Stack       = zap.Stack
	StackSkip   = zap.StackSkip
	Duration    = zap.Duration
	Durationp   = zap.Durationp
	Any         = zap.Any
)

// Initialize section
type Options struct {
	LevelEnabler Level
	Fields       []Field
	WriteSyncers []WriteSyncer
}

func InitLogger(levelEnabler Level, fields ...Field) func() error {
	return InitLoggerWithOptions(&Options{
		LevelEnabler: levelEnabler,
		Fields:       fields,
		WriteSyncers: []WriteSyncer{
			NewWriteSyncerStdout(),
		},
	})
}

func InitLoggerWithOptions(options *Options) func() error {
	// Log encoder to setup log keys
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// loading multiple write syncers from Options
	writeSyncers := zapcore.NewMultiWriteSyncer(
		options.WriteSyncers...,
	)

	// encode to json format
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	// set encoder format, writeSyncers, and lvl we want to be above
	core := zapcore.NewCore(encoder, writeSyncers, options.LevelEnabler)
	logger = zap.New(core, zap.AddCaller())

	if len(options.Fields) > 0 {
		// add addtional fields to logger if has
		logger = logger.With(options.Fields...)
	}

	// return anonymous func if success, call Sync() to turn off log resources
	return func() error {
		err := logger.Sync()

		if err != nil {
			fmt.Println("Logger Extension: Sync failed when initialize: ", err)
		}

		if file != nil {
			err = file.Close()
			if err != nil {
				fmt.Println("Logger Extension: Close file failed when initialize: ", err)
			}
		}

		if logRotation != nil {
			err = logRotation.Close()
			if err != nil {
				fmt.Println("Logger Extension: Close logRotation failed when initialize: ", err)
			}
		}

		return err
	}
}

// new writeSyncers implementations
type stdoutSync struct {
	zapcore.WriteSyncer
}

func (stdoutSync) Sync() error {
	return nil
}

type LogRotationOptions struct {
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	LocalTime  bool
}

func NewWriteSyncerStdout() WriteSyncer {
	return zapcore.AddSync(stdoutSync{os.Stdout})
}

func NewWriteSyncerFile(fileName string, logRotationOptions ...*LogRotationOptions) WriteSyncer {
	if logRotationOptions == nil || len(logRotationOptions) == 0 {
		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

		if err != nil {
			panic(fmt.Errorf("failed to create log file, err = %s", err.Error()))
		}

		return zapcore.AddSync(file)
	}

	o := logRotationOptions[0]

	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    o.MaxSize, // megabytes
		MaxBackups: o.MaxBackups,
		MaxAge:     o.MaxAge, //days
		Compress:   o.Compress,
		LocalTime:  o.LocalTime,
	})
}

// Decorate logger functions
// zap.AddCallerSkip(1) meas skip 1 step of caller info to log, equals to skip calling function, only record the caller's information (second caller)
// WithOptions will contain additional information you specify with original encoder info.
func Info(msg string, fields ...Field) {
	logger.WithOptions(zap.AddCallerSkip(1)).Info(msg, fields...)
}

func Debug(msg string, fields ...Field) {
	logger.WithOptions(zap.AddCallerSkip(1)).Debug(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	logger.WithOptions(zap.AddCallerSkip(1)).Warn(msg, fields...)
}

func Error(msg string, fields ...Field) {
	logger.WithOptions(zap.AddCallerSkip(1)).Error(msg, fields...)
}

func Panic(msg string, fields ...Field) {
	logger.WithOptions(zap.AddCallerSkip(1)).Panic(msg, fields...)
}

func ErrorStacks(msg string, fields ...Field) {
	// capture error stack by zap.Stack and named "stack_trace"
	fields = append(fields, zap.Stack("stack_trace"))
	logger.WithOptions(zap.AddCallerSkip(1)).Error(msg, fields...)
}
