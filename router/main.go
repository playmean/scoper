package router

import (
	"time"

	"github.com/playmean/scoper/common"
	"github.com/playmean/scoper/config"
	"github.com/playmean/scoper/controllers"
	"github.com/playmean/scoper/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/helmet/v2"
)

// Setup app router
func Setup(conf *config.Config, app *fiber.App) {
	app.Use(compress.New())
	app.Use(favicon.New())
	app.Use(helmet.New())

	apiGroup := app.Group("/api", basicauth.New(basicauth.Config{
		Authorizer: user.Authorizer(config.SuperUsers),
	}), controllers.MiddlewareUser)

	apiGroup.Get("/info", controllers.UserInfo)

	apiUsers := apiGroup.Group("/users", func(c *fiber.Ctx) error {
		user := c.Locals("user").(*user.User)

		if user.Role != "super" {
			return c.Status(fiber.StatusForbidden).JSON(common.Response{
				OK:    false,
				Error: "insufficient rights",
			})
		}

		return c.Next()
	})
	apiUsers.Get("/", controllers.UserList)
	apiUsers.Post("/", controllers.UserCreate)
	apiUsers.Put("/:id", controllers.UserManage)
	apiUsers.Put("/:id/reset", controllers.UserReset)

	apiProjects := apiGroup.Group("/projects")
	apiProjects.Get("/", controllers.ProjectList)
	apiProjects.Post("/", controllers.ProjectCreate)
	apiProjects.Put("/:key", controllers.ProjectManage)

	apiView := apiGroup.Group("/view/:key", controllers.MiddlewareView)
	apiView.Get("/environments", controllers.ViewEnvironments)
	apiView.Get("/tags", controllers.ViewTags)
	apiView.Get("/tags/:id", controllers.ViewTagValues)
	apiView.Get("/tracks", controllers.ViewTracks)
	apiView.Get("/tracks/:id", controllers.ViewTrack)

	app.Post("/track/:key/:type", limiter.New(limiter.Config{
		Duration: 10 * time.Second,
		Max:      10,
	}), controllers.MiddlewareTrack)
}
