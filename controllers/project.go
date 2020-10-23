package controllers

import (
	"github.com/playmean/scoper/common"
	"github.com/playmean/scoper/database"
	"github.com/playmean/scoper/project"
	"github.com/playmean/scoper/user"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

// ProjectList method
func ProjectList(c *fiber.Ctx) {
	db := database.DBConn

	list := make([]project.Project, 0)

	owner := c.Locals("user").(*user.User)

	res := db.Find(&list, "owner_id = ? OR public = ?", owner.ID, true)

	common.Answer(c, res.Error, list)
}

// ProjectCreate method
func ProjectCreate(c *fiber.Ctx) {
	db := database.DBConn

	if !common.HaveFields(c, []string{"name", "title"}) {
		return
	}

	name := c.FormValue("name")

	if !common.ValidateName(name) {
		c.Status(fiber.StatusBadRequest).JSON(common.Response{
			OK:    false,
			Error: "project name must be alphanumeric",
		})

		return
	}

	owner := c.Locals("user").(*user.User)

	keyUUID, _ := uuid.NewRandom()

	prj := project.Project{
		Key:     keyUUID.String(),
		Name:    name,
		Title:   c.FormValue("title"),
		OwnerID: owner.ID,
		Public:  false,
	}

	res := db.Create(&prj)

	common.Answer(c, res.Error, prj)
}

// ProjectManage method
func ProjectManage(c *fiber.Ctx) {
	db := database.DBConn

	if !common.HaveFields(c, []string{"name", "title", "public"}) {
		return
	}

	projectKey := c.Params("key")

	var prj project.Project

	db.First(&prj, "key = ?", projectKey)

	if prj.ID == 0 {
		c.Status(fiber.StatusNotFound).JSON(common.Response{
			OK:    false,
			Error: "project not found",
		})

		return
	}

	name := c.FormValue("name")
	title := c.FormValue("title")
	public := c.FormValue("public")

	if !common.ValidateName(name) {
		c.Status(fiber.StatusBadRequest).JSON(common.Response{
			OK:    false,
			Error: "project name must be alphanumeric",
		})

		return
	}

	prj.Name = name
	prj.Title = title
	prj.Public = public == "1"

	res := db.Save(&prj)

	common.Answer(c, res.Error, prj)
}
