package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/nsrvel/go-fiber-boilerplate/config"
	"github.com/nsrvel/go-fiber-boilerplate/internal/middleware"
	"github.com/nsrvel/go-fiber-boilerplate/internal/wrapper/handler"
	"github.com/nsrvel/go-fiber-boilerplate/internal/wrapper/repository"
	"github.com/nsrvel/go-fiber-boilerplate/internal/wrapper/usecase"
	"github.com/nsrvel/go-fiber-boilerplate/pkg/infra/db"
	"github.com/sirupsen/logrus"
)

func Run(conf *config.Config, dbList *db.DatabaseList, appLoger *logrus.Logger) {

	//* Initial Engine
	engine := html.New("./views", ".html")

	//* Initial Fiber App
	app := fiber.New(fiber.Config{
		AppName:      conf.App.Name,
		ServerHeader: "Go Fiber",
		Views:        engine,
		BodyLimit:    conf.App.BodyLimit * 1024 * 1024,
	})

	//* Initial Data Middleware
	middleware.InitMiddlewareConfig(app, dbList, conf, appLoger)

	//* General Middleware
	middleware.CORSMiddleware()
	middleware.DefaultLimitterMiddleware()
	// middleware.RecoverMiddleware()

	//* Initial Wrapper
	repo := repository.NewRepository(conf, dbList, appLoger)
	usecase := usecase.NewUsecase(repo, conf, dbList, appLoger)
	handler := handler.NewHandler(usecase, conf, appLoger)

	//* Root Endpoint
	app.Get("/", handler.General.Root.GetRoot)

	//* Api Endpoint
	// api := app.Group(conf.App.Endpoint)

	//* General Routes

	//* Core Routes

	//* CMS Routes

	//* Not found
	app.All("*", handler.General.NotFound.GetNotFound)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", conf.App.Port)))
}
