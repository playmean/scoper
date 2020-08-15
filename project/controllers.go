package project

import (
	"git.playmean.xyz/playmean/error-tracking/common"
	"git.playmean.xyz/playmean/error-tracking/database"
	"git.playmean.xyz/playmean/error-tracking/user"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

// ControllerList method
func ControllerList(c *fiber.Ctx) {
	db := database.DBConn

	list := new([]Project)

	owner := c.Locals("user").(*user.User)

	db.Find(&list, "owner_id = ?", owner.ID)

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

	keyUUID, _ := uuid.NewRandom()

	project := Project{
		Key:     keyUUID.String(),
		Name:    name,
		Title:   c.FormValue("title"),
		OwnerID: owner.ID,
	}

	db.Create(&project)

	c.JSON(common.Response{
		OK:   true,
		Data: project,
	})
}
