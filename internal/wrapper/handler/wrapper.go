package handler

import (
	"github.com/nsrvel/go-fiber-boilerplate/config"
	"github.com/nsrvel/go-fiber-boilerplate/internal/wrapper/handler/cms"
	"github.com/nsrvel/go-fiber-boilerplate/internal/wrapper/handler/core"
	"github.com/nsrvel/go-fiber-boilerplate/internal/wrapper/handler/general"
	"github.com/nsrvel/go-fiber-boilerplate/internal/wrapper/usecase"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	General general.GeneralHandler
	Core    core.CoreHandler
	CMS     cms.CMSHandler
}

func NewHandler(uc usecase.Usecase, conf *config.Config, log *logrus.Logger) Handler {
	return Handler{
		General: general.NewGeneralHandler(uc, conf, log),
		Core:    core.NewCoreHandler(uc, conf, log),
		CMS:     cms.NewCMSHandler(uc, conf, log),
	}
}
