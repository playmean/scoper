package router

import (
	"git.playmean.xyz/playmean/error-tracking/common"
	"git.playmean.xyz/playmean/error-tracking/config"
	"git.playmean.xyz/playmean/error-tracking/project"
	"git.playmean.xyz/playmean/error-tracking/track"
	"git.playmean.xyz/playmean/error-tracking/user"

	"github.com/gofiber/basicauth"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/helmet"
	"github.com/gofiber/limiter"
)

// Setup app router
func Setup(conf *config.Config, app *fiber.App) {
	app.Use(middleware.Compress())
	app.Use(middleware.Favicon())
	app.Use(helmet.New())

	apiGroup := app.Group("/api", basicauth.New(basicauth.Config{
		Authorizer: user.Authorizer(config.SuperUsers),
	}), user.Info)

	apiUsers := apiGroup.Group("/users", func(c *fiber.Ctx) {
		user := c.Locals("user").(*user.User)

		if user.Role != "super" {
			c.Status(403).JSON(common.Response{
				OK:    false,
				Error: "insufficient rights",
			})

			return
		}

		c.Next()
	})
	apiUsers.Get("/list", user.ControllerList)
	apiUsers.Get("/create", user.ControllerCreate)

	apiProjects := apiGroup.Group("/projects")
	apiProjects.Get("/list", project.ControllerList)
	apiProjects.Get("/create", project.ControllerCreate)

	app.Post("/track/:key/:type", limiter.New(limiter.Config{
		Timeout: 10,
		Max:     5,
	}), track.Middleware)
}
