package common

import (
	"regexp"

	"github.com/gofiber/fiber"
)

var regexUsername = regexp.MustCompile(`^[a-z0-9_\.]+$`)

// HaveFields in request form
func HaveFields(c *fiber.Ctx, fields []string) bool {
	output := make(map[string]string)

	for _, name := range fields {
		value := c.FormValue(name)

		if len(value) == 0 {
			c.JSON(Response{
				OK:    false,
				Error: "missing field: " + name,
			})

			return false
		}

		output[name] = value
	}

	return true
}

// ValidateName of entity
func ValidateName(name string) bool {
	return regexUsername.MatchString(name)
}
