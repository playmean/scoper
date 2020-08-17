package track

import "github.com/gofiber/fiber"

func resolveEnvironment(c *fiber.Ctx) string {
	return c.Get("X-Environment")
}
