package logger

import (
	"github.com/ic2hrmk/lokigrus"
	"github.com/nsrvel/go-fiber-boilerplate/config"
	"github.com/sirupsen/logrus"
)

func InitLoki(conf *config.Config, log *logrus.Logger) error {

	//* Logrus hook to promptail
	appLabels := make(map[string]string)
	appLabels["appName"] = conf.App.Name
	appLabels["appEnv"] = conf.App.Env

	err := InitPromtailSupport(logger, conf.Grafana.LokiURL, appLabels)
	if err != nil {
		return err
	}
	return nil
}

func InitPromtailSupport(logger *logrus.Logger, lokiAddress string, appLabels map[string]string) error {
	promtailHook, err := lokigrus.NewPromtailHook(lokiAddress, appLabels)
	if err != nil {
		return err
	}
	logger.AddHook(promtailHook)
	return nil
}
