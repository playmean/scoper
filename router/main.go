package router

import (
	"git.playmean.xyz/playmean/scoper/common"
	"git.playmean.xyz/playmean/scoper/config"
	"git.playmean.xyz/playmean/scoper/controllers"
	"git.playmean.xyz/playmean/scoper/user"

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
	}), controllers.MiddlewareUser)

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
	apiUsers.Get("/", controllers.UserList)
	apiUsers.Post("/", controllers.UserCreate)
	apiUsers.Put("/:id", controllers.UserManage)
	apiUsers.Put("/:id/reset", controllers.UserReset)

	apiProjects := apiGroup.Group("/projects")
	apiProjects.Get("/", project.ControllerList)
	apiProjects.Post("/", project.ControllerCreate)
	apiProjects.Put("/:key", project.ControllerManage)

	app.Post("/track/:key/:type", limiter.New(limiter.Config{
		Timeout: 10,
		Max:     5,
	}), track.Middleware)
}
