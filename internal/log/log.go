package log

import (
	"github.com/feditools/relay/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"reflect"
	"strings"
)

// Init the logging engine
func Init() error {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logLevel := viper.GetString(config.Keys.LogLevel)

	if logLevel != "" {
		level, err := logrus.ParseLevel(logLevel)
		if err != nil {
			return err
		}
		logrus.SetLevel(level)

		/*if level == logrus.TraceLevel {
			logrus.SetReportCaller(true)
		}*/
	}

	return nil
}

// WithPackageField creates a new logrus entry with the package name added as a field.
func WithPackageField(m interface{}) *logrus.Entry {
	return logrus.WithField("package", strings.ReplaceAll(strings.TrimPrefix(reflect.TypeOf(m).PkgPath(), "github.com/feditools/relay/"), "/", "."))
}
