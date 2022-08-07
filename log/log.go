package log

import (
	// "os"

	"time"
	"xml2pb/config"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.SetReportCaller(true)
	Log.SetFormatter(
		&formatter.Formatter{
			HideKeys:        false,
			FieldsOrder:     []string{"component", "category", "req"},
			CallerFirst:     true,
			TimestampFormat: time.RFC3339,
			NoColors:        true,
		})

	// logWriter, err := os.OpenFile(config.ConfigData.LogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	// if err != nil {
	// 	Log.Fatal(err)
	// }

	// Log.SetOutput(logWriter)
	if config.ConfigData.LogLevel == 700 {
		Log.SetLevel(logrus.DebugLevel)
	} else if config.ConfigData.LogLevel == 300 {
		Log.SetLevel(logrus.ErrorLevel)
	}
}
