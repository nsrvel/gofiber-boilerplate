package middleware

import (
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func RecoverMiddleware() {
	app := initData.App
	app.Use(recover.New())
}
