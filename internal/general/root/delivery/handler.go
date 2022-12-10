package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nsrvel/go-fiber-boilerplate/config"
	"github.com/nsrvel/go-fiber-boilerplate/internal/wrapper/usecase"
	"github.com/sirupsen/logrus"
)

type RootHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewRootHandler(uc usecase.Usecase, conf *config.Config, logger *logrus.Logger) RootHandler {
	return RootHandler{
		Usecase: uc,
		Conf:    conf,
		Log:     logger,
	}
}

func (h RootHandler) GetRoot(c *fiber.Ctx) error {
	return c.Render("start", fiber.Map{
		"Title": h.Conf.App.Name,
	})
}
