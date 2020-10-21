package track

import "github.com/gofiber/fiber"

// ResolveEnvironment for track
func ResolveEnvironment(c *fiber.Ctx) string {
	return c.Get("X-Environment")
}
