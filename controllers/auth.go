package controllers

import (
	"github.com/playmean/scoper/common"
	"github.com/playmean/scoper/user"

	"github.com/gofiber/fiber/v2"
)

// Login method
func Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	valid := user.Authorize(username, password)

	if !valid {
		return c.Status(fiber.StatusUnauthorized).JSON(common.Response{
			OK:    false,
			Error: "invalid credentials",
		})
	}

	token, err := makeToken(username)

	if err != nil {
		return common.Answer(c, err, nil)
	}

	return c.JSON(common.Response{
		OK: true,
		Data: respLogin{
			Token: token,
		},
	})
}

// RevokeToken method
func RevokeToken(c *fiber.Ctx) error {
	token, err := makeToken(c.Locals("username").(string))

	if err != nil {
		return common.Answer(c, err, nil)
	}

	return c.JSON(common.Response{
		OK: true,
		Data: respLogin{
			Token: token,
		},
	})
}
