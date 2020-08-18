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

	db.Find(&list, "owner_id = ? OR public = ?", owner.ID, true)

	c.JSON(common.Response{
		OK:   true,
		Data: list,
	})
}

// ControllerCreate method
func ControllerCreate(c *fiber.Ctx) {
	db := database.DBConn

	if !common.HaveFields(c, []string{"name", "title"}) {
		return
	}

	name := c.FormValue("name")

	if !common.ValidateName(name) {
		c.Status(400).JSON(common.Response{
			OK:    false,
			Error: "project name must be alphanumeric",
		})

		return
	}

	owner := c.Locals("user").(*user.User)

	keyUUID, _ := uuid.NewRandom()

	prj := Project{
		Key:     keyUUID.String(),
		Name:    name,
		Title:   c.FormValue("title"),
		OwnerID: owner.ID,
		Public:  false,
	}

	db.Create(&prj)

	c.JSON(common.Response{
		OK:   true,
		Data: prj,
	})
}

// ControllerManage method
func ControllerManage(c *fiber.Ctx) {
	db := database.DBConn

	if !common.HaveFields(c, []string{"name", "title", "public"}) {
		return
	}

	projectKey := c.Params("key")

	var prj Project

	db.First(&prj, "key = ?", projectKey)

	if prj.ID == 0 {
		c.Status(404).JSON(common.Response{
			OK:    false,
			Error: "project not found",
		})

		return
	}

	name := c.FormValue("name")
	title := c.FormValue("title")
	public := c.FormValue("public")

	if !common.ValidateName(name) {
		c.Status(400).JSON(common.Response{
			OK:    false,
			Error: "project name must be alphanumeric",
		})

		return
	}

	prj.Name = name
	prj.Title = title
	prj.Public = public == "1"

	db.Save(&prj)

	c.JSON(common.Response{
		OK:   true,
		Data: prj,
	})
}
