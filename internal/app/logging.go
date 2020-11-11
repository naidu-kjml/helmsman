package app

import (
	"net/url"
	"os"

	"github.com/apsdehal/go-logger"
)

type Logger struct {
	SlackWebhook string
	baseLogger   *logger.Logger
}

var log = &Logger{}

func (l *Logger) Info(message string) {
	l.baseLogger.Info(message)
}

func (l *Logger) Debug(message string) {
	if flags.debug {
		l.baseLogger.Debug(message)
	}
}

func (l *Logger) Verbose(message string) {
	if flags.verbose {
		l.baseLogger.Info(message)
	}
}

func (l *Logger) Error(message string) {
	if _, err := url.ParseRequestURI(l.SlackWebhook); err == nil {
		notifySlack(message, l.SlackWebhook, true, flags.apply)
	}
	l.baseLogger.Error(message)
}

func (l *Logger) Errorf(message string, args ...interface{}) {
	l.baseLogger.Errorf(message, args...)
}

func (l *Logger) Warning(message string) {
	l.baseLogger.Warning(message)
}

func (l *Logger) Notice(message string) {
	l.baseLogger.Notice(message)
}

func (l *Logger) Critical(message string) {
	if _, err := url.ParseRequestURI(l.SlackWebhook); err == nil {
		notifySlack(message, l.SlackWebhook, true, flags.apply)
	}
	l.baseLogger.Critical(message)
}

func (l *Logger) Fatal(message string) {
	if _, err := url.ParseRequestURI(l.SlackWebhook); err == nil {
		notifySlack(message, l.SlackWebhook, true, flags.apply)
	}
	l.baseLogger.Fatal(message)
}

func initLogs(verbose bool, noColors bool) {
	logger.SetDefaultFormat("%{time:2006-01-02 15:04:05} %{level}: %{message}")
	logLevel := logger.InfoLevel
	if verbose {
		logLevel = logger.DebugLevel
	}
	colors := 1
	if noColors {
		colors = 0
	}
	log.baseLogger, _ = logger.New("logger", colors, os.Stdout, logLevel)
}
