package router

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func checkAuthScheme(c *fiber.Ctx, authScheme string) bool {
	authHeader := c.Get(fiber.HeaderAuthorization)

	headerLen := len(authHeader)
	schemeLen := len(authScheme)

	return headerLen > schemeLen+1 && strings.ToLower(authHeader[:schemeLen]) == authScheme
}
