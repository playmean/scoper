package controllers

import (
	"github.com/playmean/scoper/common"
	"github.com/playmean/scoper/database"
	"github.com/playmean/scoper/project"
	"github.com/playmean/scoper/user"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// ProjectList method
func ProjectList(c *fiber.Ctx) error {
	db := database.DBConn

	list := make([]project.Project, 0)

	owner := c.Locals("user").(*user.User)

	res := db.Find(&list, "owner_id = ? OR public = ?", owner.ID, true)

	return common.Answer(c, res.Error, list)
}

// ProjectCreate method
func ProjectCreate(c *fiber.Ctx) error {
	db := database.DBConn

	if !common.HaveFields(c, []string{"name", "title"}) {
		return c.Next()
	}

	name := c.FormValue("name")

	if !common.ValidateName(name) {
		return c.Status(fiber.StatusBadRequest).JSON(common.Response{
			OK:    false,
			Error: "project name must be alphanumeric",
		})
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

	return common.Answer(c, res.Error, prj)
}

// ProjectManage method
func ProjectManage(c *fiber.Ctx) error {
	db := database.DBConn

	if !common.HaveFields(c, []string{"name", "title", "public"}) {
		return c.Next()
	}

	projectKey := c.Params("key")

	var prj project.Project

	db.First(&prj, "key = ?", projectKey)

	if prj.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(common.Response{
			OK:    false,
			Error: "project not found",
		})
	}

	name := c.FormValue("name")
	title := c.FormValue("title")
	public := c.FormValue("public")

	if !common.ValidateName(name) {
		return c.Status(fiber.StatusBadRequest).JSON(common.Response{
			OK:    false,
			Error: "project name must be alphanumeric",
		})
	}

	prj.Name = name
	prj.Title = title
	prj.Public = public == "1"

	res := db.Save(&prj)

	return common.Answer(c, res.Error, prj)
}
