package controllers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/playmean/scoper/common"
	"github.com/playmean/scoper/user"
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

	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString(common.SigningKey)
	if err != nil {
		return common.Answer(c, err, nil)
	}

	return c.JSON(common.Response{
		OK: true,
		Data: respLogin{
			Token: t,
		},
	})
}
