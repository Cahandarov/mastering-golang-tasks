package logger

import (
	"github.com/sirupsen/logrus"
)

func Init() error {
	lvl := logrus.InfoLevel
	logrus.SetLevel(lvl)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		DisableColors: false,
		ForceColors:   true,
	})

	logrus.Info("Log level is ", lvl)
	return nil
}
