package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nsrvel/go-fiber-boilerplate/config"
	"github.com/nsrvel/go-fiber-boilerplate/pkg/infra/db"
	"github.com/sirupsen/logrus"
)

//* Init
var initData MiddlewareData

type MiddlewareData struct {
	App    *fiber.App
	DBList *db.DatabaseList
	Conf   *config.Config
	Log    *logrus.Logger
}

func InitMiddlewareConfig(app *fiber.App, dbList *db.DatabaseList, conf *config.Config, log *logrus.Logger) {
	initData = MiddlewareData{
		App:    app,
		DBList: dbList,
		Conf:   conf,
		Log:    log,
	}
}
