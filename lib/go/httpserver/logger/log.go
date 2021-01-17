package logger

import (
	"context"
	"errors"
	"time"

	stackdriver "github.com/TV4/logrus-stackdriver-formatter"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// Level for logger
type Level uint32

// String form of a Level
func (l Level) String() string {
	return logrus.Level(l).String()
}

// SetLevel is used to change the level of logging globally
var SetLevel = func(level Level) {
	logrus.SetLevel(logrus.Level(level))
}

// GetLevel is used to get the current level of logging
var GetLevel = func() Level {
	return Level(logrus.GetLevel())
}

// DebugLevel is a threshold level
const DebugLevel Level = Level(logrus.DebugLevel)

// InfoLevel is a threshold level
const InfoLevel Level = Level(logrus.InfoLevel)

// WarnLevel is a threshold level
const WarnLevel Level = Level(logrus.WarnLevel)

// ErrorLevel is a threshold level
const ErrorLevel Level = Level(logrus.ErrorLevel)

// FatalLevel is a threshold level
const FatalLevel Level = Level(logrus.FatalLevel)

// Interface for the logger
type Interface = *logrus.Entry

// InterfaceWithContext for the logger with context in arguments
type InterfaceWithContext interface {
	LogMode(gormlogger.LogLevel) gormlogger.Interface
	Info(ctx context.Context, s string, args ...interface{})
	Warn(ctx context.Context, s string, args ...interface{})
	Error(ctx context.Context, s string, args ...interface{})
	Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error)
}

// New logger
func New(moduleName string, usingStackdriver bool) Interface {
	logrus.SetLevel(logrus.InfoLevel)
	if usingStackdriver {
		logrus.SetFormatter(stackdriver.NewFormatter(
			stackdriver.WithService(moduleName),
		))
	}
	return logrus.WithFields(logrus.Fields{
		"module": moduleName,
	})
}

// instanceWithContext is the logger with conext in arguments
type instanceWithContext struct {
	log                   Interface
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
}

// ForGORM creates a new logger with context in arguments
func ForGORM(log Interface) InterfaceWithContext {
	return &instanceWithContext{
		log:                   log,
		SkipErrRecordNotFound: true,
	}
}

// LogMode is a stub for changing the log level
func (l *instanceWithContext) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

// Info log with context
func (l *instanceWithContext) Info(ctx context.Context, s string, args ...interface{}) {
	l.log.WithContext(ctx).Infof(s, args...)
}

// Warn log with context
func (l *instanceWithContext) Warn(ctx context.Context, s string, args ...interface{}) {
	l.log.WithContext(ctx).Warnf(s, args...)
}

// Error log with context
func (l *instanceWithContext) Error(ctx context.Context, s string, args ...interface{}) {
	l.log.WithContext(ctx).Errorf(s, args...)
}

// Trace log with context
func (l *instanceWithContext) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := logrus.Fields{}
	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		fields[logrus.ErrorKey] = err
		l.log.WithContext(ctx).WithFields(fields).Errorf("%s [%s]", sql, elapsed)
		return
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		l.log.WithContext(ctx).WithFields(fields).Warnf("%s [%s]", sql, elapsed)
		return
	}

	l.log.WithContext(ctx).WithFields(fields).Debugf("%s [%s]", sql, elapsed)
}
