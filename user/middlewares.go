package user

import (
	"git.playmean.xyz/playmean/scoper/database"

	"github.com/gofiber/fiber"
)

// Info about user
func Info(c *fiber.Ctx) {
	db := database.DBConn

	user := User{
		Username: c.Locals("username").(string),
		Role:     "super",
	}

	db.Where("username = ?", user.Username).Take(&user)

	c.Locals("user", &user)

	c.Next()
}
