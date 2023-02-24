// Package logger provides a wrapper around zerolog, a zero allocation, JSON logging library.
package torolog

import (
	"io"

	"github.com/rs/zerolog"
)

// fieldsKey is used as the key for the fields in the log message.
var fieldsKey = "data"

// Level defines the log levels that are used in the logger.
type Level int

// Log level used in the logger.
const (
	TraceLevel Level = iota - 1
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
	NoLevel
)

// Field is used to specify a key-value pair in the log message.
type Field struct {
	Key   string
	Value interface{}
}

// MarshalZerologObject is a method that implements the zerolog.ObjectMarshaler interface.
func (f Field) MarshalZerologObject(e *zerolog.Event) {
	e.Interface(f.Key, f.Value)
}

// Fields is a slice of Field.
type Fields []Field

// MarshalZerologArray is a method that implements the zerolog.ArrayMarshaler interface.
func (ff Fields) MarshalZerologArray(a *zerolog.Array) {
	for _, f := range ff {
		a.Object(f)
	}
}

// Logger is a wrapper around zerolog.Logger.
type Logger struct {
	log   *zerolog.Logger
	level Level
}

// New creates a enw Logger instance with the given io.Writer output.
func New(output io.Writer) *Logger {
	logger := zerolog.New(output)
	return &Logger{&logger, NoLevel}
}

// NewWithLevel creates a new Logger instance with the given io.Writer output and log level.
func NewWithLevel(output io.Writer, level Level) *Logger {
	var logger zerolog.Logger
	if level < -1 || level > 6 {
		return New(output)
	}
	logger = zerolog.New(output).Level(zerolog.Level(level))
	return &Logger{&logger, level}
}

// Log logs a message
func (l *Logger) Log(msg string, fields Fields) {
	l.log.Log().Array(fieldsKey, fields).Msg(msg)
}

// Trace logs a message with log level Trace.
func (l *Logger) Trace(msg string, err error, fields Fields) {
	l.log.Trace().Array(fieldsKey, fields).Err(err).Msg(msg)
}

// Debug logs a message with log level Debug.
func (l *Logger) Debug(msg string, fields Fields) {
	l.log.Debug().Array(fieldsKey, fields).Msg(msg)
}

// Info logs a message with log level Info.
func (l *Logger) Info(msg string, fields Fields) {
	l.log.Info().Array(fieldsKey, fields).Msg(msg)
}

// Warn logs a message with log level Warn.
func (l *Logger) Warn(msg string, fields Fields) {
	l.log.Warn().Array(fieldsKey, fields).Msg(msg)
}

// Error logs a message with log level Error.
func (l *Logger) Error(msg string, err error, fields Fields) {
	l.log.Error().Array(fieldsKey, fields).Err(err).Msg(msg)
}

// Fatal logs a message with log level Fatal.
func (l *Logger) Fatal(msg string, err error, fields Fields) {
	l.log.Fatal().Array(fieldsKey, fields).Err(err).Msg(msg)
}

// Panic logs a message with log level Panic.
func (l *Logger) Panic(msg string, err error, fields Fields) {
	l.log.Panic().Array(fieldsKey, fields).Err(err).Msg(msg)
}

// GetLevel returns the logger log level.
func (l *Logger) GetLevel() Level {
	return l.level
}
