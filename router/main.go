package router

import (
	"crypto/rand"
	"crypto/rsa"
	"time"

	"github.com/playmean/scoper/common"
	"github.com/playmean/scoper/config"
	"github.com/playmean/scoper/controllers"
	"github.com/playmean/scoper/logger"
	"github.com/playmean/scoper/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/helmet/v2"

	jwtware "github.com/gofiber/jwt/v2"
)

// Setup app router
func Setup(conf *config.Config, app *fiber.App) {
	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		logger.Fatal("KEY", "%v", err)
	}

	common.SigningKey = rsaKey

	app.Use(compress.New())
	app.Use(favicon.New())
	app.Use(helmet.New())

	apiGroup := app.Group("/api")
	apiGroup.Post("/login", controllers.Login)
	apiGroup.Use(jwtware.New(jwtware.Config{
		SigningMethod: "RS256",
		SigningKey:    common.SigningKey.Public(),
		ContextKey:    "jwt",
		Filter: func(c *fiber.Ctx) bool {
			return checkAuthScheme(c, "basic")
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(common.Response{
				OK:    false,
				Error: "invalid or expired token",
			})
		},
	}))
	apiGroup.Use(basicauth.New(basicauth.Config{
		Authorizer: user.Authorize,
		Next: func(c *fiber.Ctx) bool {
			return c.Locals("jwt") != nil
		},
	}))
	apiGroup.Use(controllers.MiddlewareUser)

	apiGroup.Get("/info", controllers.UserInfo)
	apiGroup.Put("/revoke", controllers.RevokeToken)

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
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(common.Response{
				OK:    false,
				Error: "too many requests",
			})
		},
	}), controllers.MiddlewareTrack)
}
