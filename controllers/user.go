package controllers

import (
	"github.com/playmean/scoper/common"
	"github.com/playmean/scoper/database"
	"github.com/playmean/scoper/generator"
	"github.com/playmean/scoper/user"

	"github.com/gofiber/fiber"
)

// MiddlewareUser method
func MiddlewareUser(c *fiber.Ctx) {
	db := database.DBConn

	u := user.User{
		Username: c.Locals("username").(string),
		Role:     "super",
	}

	db.Where("username = ?", u.Username).Take(&u)

	c.Locals("user", &u)

	c.Next()
}

// UserInfo method
func UserInfo(c *fiber.Ctx) {
	u := c.Locals("user").(*user.User)

	c.JSON(common.Response{
		OK: true,
		Data: respUserInfo{
			ID: u.ID,

			Username: u.Username,
			FullName: u.FullName,
			Role:     u.Role,

			CreatedAt: u.CreatedAt,
		},
	})
}

// UserList method
func UserList(c *fiber.Ctx) {
	db := database.DBConn

	list := make([]user.User, 0)

	db.Find(&list)

	c.JSON(common.Response{
		OK:   true,
		Data: list,
	})
}

// UserCreate method
func UserCreate(c *fiber.Ctx) {
	db := database.DBConn

	if !common.HaveFields(c, []string{"username", "role"}) {
		return
	}

	username := c.FormValue("username")

	if !common.ValidateName(username) {
		c.Status(fiber.StatusBadRequest).JSON(common.Response{
			OK:    false,
			Error: "username must be alphanumeric",
		})

		return
	}

	var u user.User

	db.First(&u, "username = ?", username)

	if u.ID > 0 {
		c.Status(fiber.StatusConflict).JSON(common.Response{
			OK:    false,
			Error: "username already taken",
		})

		return
	}

	password := generator.Password(12)

	u = user.User{
		Username:     username,
		FullName:     c.FormValue("fullname"),
		Password:     password,
		PasswordHash: user.HashPassword(password),
		Role:         c.FormValue("role"),
	}

	res := db.Save(&u)

	common.Answer(c, res.Error, u)
}

// UserManage method
func UserManage(c *fiber.Ctx) {
	db := database.DBConn

	if !common.HaveFields(c, []string{"role"}) {
		return
	}

	userID := c.Params("id")

	var u user.User

	db.First(&u, userID)

	if u.ID == 0 {
		c.Status(fiber.StatusNotFound).JSON(common.Response{
			OK:    false,
			Error: "user not found",
		})

		return
	}

	fullname := c.FormValue("fullname")
	role := c.FormValue("role")

	u.FullName = fullname
	u.Role = role

	res := db.Save(&u)

	common.Answer(c, res.Error, u)
}

// UserReset method
func UserReset(c *fiber.Ctx) {
	db := database.DBConn

	userID := c.Params("id")

	var u user.User

	db.First(&u, userID)

	if u.ID == 0 {
		c.Status(fiber.StatusNotFound).JSON(common.Response{
			OK:    false,
			Error: "user not found",
		})

		return
	}

	password := generator.Password(12)

	u.Password = password
	u.PasswordHash = user.HashPassword(password)

	res := db.Save(&u)

	common.Answer(c, res.Error, u)
}
