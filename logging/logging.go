package logging

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

func init() {
	logrus.SetFormatter(new(LoggerFormatter))
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetOutput(os.Stdout)
}

var AppName string

/**
 * Configure logging
 */
func Configure(appName, output, level string) {
	AppName = appName
	if output == "" || output == "stdout" {
		logrus.SetOutput(os.Stdout)
	} else if output == "stderr" {
		logrus.SetOutput(os.Stderr)
	} else {
		f, err := os.OpenFile(output, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
		if err != nil {
			logrus.Fatal(err)
		}
		mw := io.MultiWriter(os.Stdout, f)
		logrus.SetOutput(mw)
	}

	if level == "" {
		return
	}

	if level, err := logrus.ParseLevel(level); err != nil {
		logrus.Fatal("Unknown loglevel ", level)
	} else {
		logrus.SetLevel(level)
	}
}

/**
 * Our custom formatter
 */
type LoggerFormatter struct{}

/**
 * Format entry
 */
func (f *LoggerFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	b := &bytes.Buffer{}
	name, ok := entry.Data["name"]
	app, _ := entry.Data["app"]
	if !ok {
		name = "default"
	}
	fmt.Fprintf(b, "[%s] %s [%-5.5s] (%s): %s\n", app, entry.Time.Format("2006-01-02 15:04:05"), strings.ToUpper(entry.Level.String()), name, entry.Message)
	return b.Bytes(), nil
}

/**
 * Add logger name as field var
 */
func For(name string) *logrus.Entry {
	return logrus.WithField("name", name).WithField("app", AppName)
}

func Alarm(log logrus.Entry) *logrus.Entry {
	return log.WithField("alarm", true)
}
