package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	// do something here to set environment depending on an environment variable
	// or command-line flag
	Environment := os.Getenv("env")
	if Environment == "" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		// The TextFormatter is default, you don't actually have to do this.
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

}

func main() {
	logrus.SetLevel(logrus.TraceLevel)

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	logrus.Fatal("fatal msg")
	logrus.Panic("panic msg")
}
