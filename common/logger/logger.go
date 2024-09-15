package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
)

var MyLogger *logrus.Logger

func init() {
	MyLogger = logrus.New()
	MyLogger.SetOutput(os.Stdout)
	MyLogger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceQuote:      true,
	})
}

func LogError(err error) {
	MyLogger.Error(err.Error())
}

func LogErrorStr(err string) {
	MyLogger.Error(err)
}

func LogInfo(info string) {
	MyLogger.Info(info)
}

func Info(args ...interface{}) {
	MyLogger.Info(args...)
}

func Error(args ...interface{}) {
	MyLogger.Error(args...)
}

func Infof(format string, args ...interface{}) {
	MyLogger.Infof(format, args...)
}

func LogInfof(format string, args ...interface{}) {
	MyLogger.Infof(format, args...)
}

func LogWithFields(fields logrus.Fields) *logrus.Entry {
	return MyLogger.WithFields(fields)
}

func Fatal(err error) {
	MyLogger.Fatal(err.Error())
	os.Exit(1)
}

func Fatalf(format string, args ...interface{}) {
	MyLogger.Fatalf(format, args...)
	os.Exit(1)
}

func WithContext(ctx context.Context) *logrus.Entry {

	return MyLogger.WithContext(ctx)
}
