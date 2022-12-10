package general

import (
	"github.com/nsrvel/go-fiber-boilerplate/config"
	notfound "github.com/nsrvel/go-fiber-boilerplate/internal/general/notfound/delivery"
	root "github.com/nsrvel/go-fiber-boilerplate/internal/general/root/delivery"
	"github.com/nsrvel/go-fiber-boilerplate/internal/wrapper/usecase"
	"github.com/sirupsen/logrus"
)

type GeneralHandler struct {
	Root     root.RootHandler
	NotFound notfound.NotFoundHandler
}

func NewGeneralHandler(uc usecase.Usecase, conf *config.Config, log *logrus.Logger) GeneralHandler {
	return GeneralHandler{
		Root:     root.NewRootHandler(uc, conf, log),
		NotFound: notfound.NewNotFoundHandler(uc, conf, log),
	}
}
