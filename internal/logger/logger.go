package logger

import (
	"github.com/Vellvill/frames/internal/config"
	"github.com/sirupsen/logrus"
	"sync"
)

var once sync.Once
var Logger *logrus.Logger

func init() {
	once.Do(func() {
		Logger = logrus.New()
		Logger.SetLevel(logrus.Level(config.GetValue("log_level").Int()))
		Logger.Debug("logger started")
	})
}
