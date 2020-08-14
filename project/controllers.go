package project

import (
	"git.playmean.xyz/playmean/error-tracking/common"
	"git.playmean.xyz/playmean/error-tracking/database"
	"git.playmean.xyz/playmean/error-tracking/user"

	"github.com/gofiber/fiber"
)

// ControllerList method
func ControllerList(c *fiber.Ctx) {
	db := database.DBConn

	list := new([]Project)

	db.Find(&list)

	c.JSON(common.Response{
		OK:   true,
		Data: list,
	})
}

// ControllerCreate method
func ControllerCreate(c *fiber.Ctx) {
	if !common.HaveFields(c, []string{"name", "title"}) {
		return
	}

	name := c.FormValue("name")

	if !common.ValidateName(name) {
		c.JSON(common.Response{
			OK:    false,
			Error: "project name must be alphanumeric",
		})

		return
	}

	db := database.DBConn

	owner := c.Locals("user").(*user.User)

	project := Project{
		Name:  name,
		Title: c.FormValue("title"),
		Owner: *owner,
	}

	db.Create(&project)

	c.JSON(common.Response{
		OK:   true,
		Data: project,
	})
}
