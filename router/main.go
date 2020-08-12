package router

import (
	"git.playmean.xyz/playmean/error-tracking/config"
	"git.playmean.xyz/playmean/error-tracking/track"
	"git.playmean.xyz/playmean/error-tracking/user"

	"github.com/gofiber/basicauth"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/helmet"
)

// Setup app router
func Setup(conf *config.Config, app *fiber.App) {
	app.Use(middleware.Compress())
	app.Use(middleware.Favicon())
	app.Use(helmet.New())

	apiGroup := app.Group("/api", basicauth.New(basicauth.Config{
		Authorizer: user.Authorizer(map[string]string{
			"super": conf.Password,
		}),
	}))

	apiGroup.Group("/users", func(c *fiber.Ctx) {
		if c.Locals("role").(string) != "admin" {
			c.Status(403).JSON(response{
				OK:    false,
				Error: "insufficient rights",
			})

			return
		}
	})

	trackGroup := app.Group("/track/:key", track.Middleware)
	trackGroup.Post("/error", track.Error)
	trackGroup.Post("/log", track.Log)
}
