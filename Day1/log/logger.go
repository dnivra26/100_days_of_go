package log

import (
	"os"

	"github.com/Sirupsen/logrus"
)

func InitLogDevConfig() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

func InitLogProdConfig() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.WarnLevel)
}

func Info(msg string) {
	logrus.Info(msg)
}

func Debug(msg string) {
	logrus.Debug(msg)
}

func Warn(msg string) {
	logrus.Warn(msg)
}

func Error(msg string) {
	logrus.Error(msg)
}

func Fatal(msg string) {
	logrus.Fatal(msg)
}

func Panic(msg string) {
	logrus.Panic(msg)
}

func Init(environment string) {
	switch environment {
	case "DEV":
		InitLogDevConfig()
	case "PRODUCTION":
		InitLogProdConfig()
	}
	logrus.Debugf("Environment : %s", environment)
}
