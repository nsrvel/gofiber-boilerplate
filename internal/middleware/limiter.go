package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

//* Limiter middleware
func DefaultLimitterMiddleware() {
	app := initData.App
	app.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You have requested too many in a single time-frame! Please wait another minute!",
			})
		},
	}))
}

// func LimiterLoginMiddleware(app *fiber.App, handler handler.Handler, conf *config.Config, log *logrus.Logger) {

// 	app.Use("/api/v1/auth/login", limiter.New(limiter.Config{
// 		Max:        5,
// 		Expiration: 1 * time.Minute,
// 		KeyGenerator: func(c *fiber.Ctx) string {
// 			return c.IP()
// 		},
// 		LimitReached: func(c *fiber.Ctx) error {
// 			init := exception.InitException(c, conf, log)

// 			return exception.CreateResponse_Log(init, fiber.StatusBadRequest, "Too many failed login attempts, please try again 1 minute later ", "Terlalu banyak upaya login yang gagal, silahkan coba lagi 1 menit kemudian", nil)
// 		},
// 		SkipFailedRequests:     false,
// 		SkipSuccessfulRequests: true,
// 		LimiterMiddleware:      limiter.FixedWindow{},
// 	}))
// }

// func LimiterChangePwd(app *fiber.App, handler handler.Handler, conf *config.Config, log *logrus.Logger) {

// 	app.Use("/api/v1/user/changepassword", limiter.New(limiter.Config{
// 		Max:        5,
// 		Expiration: 1 * time.Minute,
// 		KeyGenerator: func(c *fiber.Ctx) string {
// 			return c.IP()
// 		},
// 		LimitReached: func(c *fiber.Ctx) error {
// 			init := exception.InitException(c, conf, log)

// 			return exception.CreateResponse_Log(init, fiber.StatusBadRequest, "Too many failed attempts, please try again 1 minute later ", "Terlalu banyak upaya yang gagal, silahkan coba lagi 1 menit kemudian", nil)
// 		},
// 		SkipFailedRequests:     false,
// 		SkipSuccessfulRequests: true,
// 		LimiterMiddleware:      limiter.FixedWindow{},
// 	}))
// }
