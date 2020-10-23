package track

import "github.com/gofiber/fiber/v2"

// ResolveEnvironment for track
func ResolveEnvironment(c *fiber.Ctx) string {
	return c.Get("X-Environment")
}
