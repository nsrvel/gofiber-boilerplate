package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORSMiddleware() {

	app := initData.App
	conf := initData.Conf

	//* CORS
	CORS := cors.Config{}
	CORS.AllowHeaders = conf.Routes.Headers
	CORS.AllowMethods = conf.Routes.Methods
	if conf.Routes.Origins.IsDefault {
		CORS.AllowOrigins = "*"
	} else {
		CORS.AllowOrigins = fmt.Sprintf("%s, %s, %s",
			conf.Routes.Origins.FeLocal,
			conf.Routes.Origins.FeDev,
			conf.Routes.Origins.FeProd,
		)
	}
	app.Use(cors.New(CORS))
}
