package cms

import (
	"github.com/nsrvel/go-fiber-boilerplate/config"
	"github.com/nsrvel/go-fiber-boilerplate/internal/wrapper/usecase"
	"github.com/sirupsen/logrus"
)

type CMSHandler struct {
}

func NewCMSHandler(uc usecase.Usecase, conf *config.Config, log *logrus.Logger) CMSHandler {
	return CMSHandler{}
}
