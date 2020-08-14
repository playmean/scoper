package user

import (
	"git.playmean.xyz/playmean/error-tracking/common"
	"git.playmean.xyz/playmean/error-tracking/database"

	"github.com/gofiber/fiber"
)

// ControllerList method
func ControllerList(c *fiber.Ctx) {
	db := database.DBConn

	list := new([]User)

	db.Find(&list)

	c.JSON(common.Response{
		OK:   true,
		Data: list,
	})
}

// ControllerCreate method
func ControllerCreate(c *fiber.Ctx) {
	if !common.HaveFields(c, []string{"username", "password", "role"}) {
		return
	}

	username := c.FormValue("username")

	if !common.ValidateName(username) {
		c.JSON(common.Response{
			OK:    false,
			Error: "username must be alphanumeric",
		})

		return
	}

	db := database.DBConn

	var user User

	db.Where("username = ?", username).First(&user)

	if user.ID > 0 {
		c.JSON(common.Response{
			OK:    false,
			Error: "username already taken",
		})

		return
	}

	user = User{
		Username:     username,
		PasswordHash: hashPassword(c.FormValue("password")),
		Role:         c.FormValue("role"),
	}

	db.Save(&user)

	c.JSON(common.Response{
		OK:   true,
		Data: user,
	})
}
