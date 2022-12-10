package delivery

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nsrvel/go-fiber-boilerplate/config"
	"github.com/nsrvel/go-fiber-boilerplate/internal/wrapper/usecase"
	"github.com/nsrvel/go-fiber-boilerplate/pkg/exception"
	"github.com/sirupsen/logrus"
)

type NotFoundHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewNotFoundHandler(uc usecase.Usecase, conf *config.Config, logger *logrus.Logger) NotFoundHandler {
	return NotFoundHandler{
		Usecase: uc,
		Conf:    conf,
		Log:     logger,
	}
}

func (h NotFoundHandler) GetNotFound(c *fiber.Ctx) error {

	init := exception.InitException(c, h.Conf, h.Log)

	errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

	return exception.CreateResponse(init, fiber.StatusNotFound, errorMessage, errorMessage, nil)
}
