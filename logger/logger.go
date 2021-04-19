package logger

import (
	"os"

	"github.com/Eldius/things-manager-go/config"
	"github.com/sirupsen/logrus"
)

var appLogger *logrus.Entry

func setup() {
	if appLogger == nil {
		hostname, _ := os.Hostname()
		var standardFields = logrus.Fields{
			"hostname": hostname,
			"appname":  "controle-oi-cadastro-cartao-backend",
		}

		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetReportCaller(true)
		format := config.GetLoggerFormat()
		if format == "text" {
			logrus.SetFormatter(&logrus.TextFormatter{})
		} else {
			logrus.SetFormatter(&logrus.JSONFormatter{})
		}
		appLogger = logrus.WithFields(standardFields)
	}
}

/*
Logger returns the app logger
*/
func Logger() *logrus.Entry {
	if appLogger == nil {
		setup()
	}
	return appLogger
}
