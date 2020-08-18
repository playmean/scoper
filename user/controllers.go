package user

import (
	"git.playmean.xyz/playmean/error-tracking/common"
	"git.playmean.xyz/playmean/error-tracking/database"
	"git.playmean.xyz/playmean/error-tracking/generator"

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
	db := database.DBConn

	if !common.HaveFields(c, []string{"username", "role"}) {
		return
	}

	username := c.FormValue("username")

	if !common.ValidateName(username) {
		c.Status(400).JSON(common.Response{
			OK:    false,
			Error: "username must be alphanumeric",
		})

		return
	}

	var user User

	db.First(&user, "username = ?", username)

	if user.ID > 0 {
		c.Status(409).JSON(common.Response{
			OK:    false,
			Error: "username already taken",
		})

		return
	}

	password := generator.Password(12)

	user = User{
		Username:     username,
		FullName:     c.FormValue("fullname"),
		Password:     password,
		PasswordHash: hashPassword(password),
		Role:         c.FormValue("role"),
	}

	db.Save(&user)

	c.JSON(common.Response{
		OK:   true,
		Data: user,
	})
}

// ControllerManage method
func ControllerManage(c *fiber.Ctx) {
	db := database.DBConn

	if !common.HaveFields(c, []string{"role"}) {
		return
	}

	userID := c.Params("id")

	var user User

	db.First(&user, userID)

	if user.ID == 0 {
		c.Status(404).JSON(common.Response{
			OK:    false,
			Error: "user not found",
		})

		return
	}

	fullname := c.FormValue("fullname")
	role := c.FormValue("role")

	user.FullName = fullname
	user.Role = role

	db.Save(&user)

	c.JSON(common.Response{
		OK:   true,
		Data: user,
	})
}

// ControllerReset method
func ControllerReset(c *fiber.Ctx) {
	db := database.DBConn

	userID := c.Params("id")

	var user User

	db.First(&user, userID)

	if user.ID == 0 {
		c.Status(404).JSON(common.Response{
			OK:    false,
			Error: "user not found",
		})

		return
	}

	password := generator.Password(12)

	user.Password = password
	user.PasswordHash = hashPassword(password)

	db.Save(&user)

	c.JSON(common.Response{
		OK:   true,
		Data: user,
	})
}
