package log

import "github.com/sirupsen/logrus"

var defaultLog = logrus.New()

func init() {
	defaultLog.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func GetDefaultLog() *logrus.Logger {
	return defaultLog
}
