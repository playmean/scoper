package track

import (
	"encoding/json"

	"git.playmean.xyz/playmean/error-tracking/logger"

	"github.com/gofiber/fiber"
)

// Middleware for tracking APIs
func Middleware(c *fiber.Ctx) {
	projectKey := c.Params("key")

	if projectKey == "" {
		c.JSON(response{
			OK:    false,
			Error: "project key not specified",
		})

		return
	}

	c.Next()
}

// Error tracking controller
func Error(c *fiber.Ctx) {
	projectKey := c.Params("key")

	var body reportPacket

	c.BodyParser(&body)

	formatted, _ := json.MarshalIndent(body, "", "    ")

	logger.Log("TRACK:ERROR", projectKey, string(formatted))

	c.JSON(response{
		OK: true,
	})
}

// Log tracking controller
func Log(c *fiber.Ctx) {
	projectKey := c.Params("key")

	var body interface{}

	c.BodyParser(&body)

	formatted, _ := json.MarshalIndent(body, "", "    ")

	logger.Log("TRACK:LOG", projectKey, string(formatted))

	c.JSON(response{
		OK: true,
	})
}
