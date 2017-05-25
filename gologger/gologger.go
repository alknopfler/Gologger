package gologger

import (

"github.com/sirupsen/logrus"

	"io"
)

func Init(o io.Writer, level logrus.Level) {
	logrus.SetOutput(o)
	logrus.SetLevel(level)
}

func Print(level string, errornum int, desc string, filename string) {
	fields := logrus.Fields{
		"Go-file":     filename,
		"errorNumber": errornum,
	}
	switch level {
	case "INFO":
		logrus.WithFields(fields).Info(desc)
	case "WARN":
		logrus.WithFields(fields).Warn(desc)
	case "ERROR":
		logrus.WithFields(fields).Error(desc)
	case "PANIC":
		logrus.WithFields(fields).Panic(desc)
	case "FATAL":
		logrus.WithFields(fields).Fatal(desc)
	default:
		logrus.WithFields(fields).Debug(desc)
	}
}
