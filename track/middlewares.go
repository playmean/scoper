package track

import (
	"git.playmean.xyz/playmean/error-tracking/common"
	"git.playmean.xyz/playmean/error-tracking/database"
	"git.playmean.xyz/playmean/error-tracking/project"

	"github.com/gofiber/fiber"
)

// Middleware for tracking APIs
func Middleware(c *fiber.Ctx) {
	db := database.DBConn

	var prj project.Project

	projectKey := c.Params("key")
	trackType := c.Params("type")

	db.First(&prj, "key = ?", projectKey)

	if prj.ID == 0 {
		c.JSON(common.Response{
			OK:    false,
			Error: "project not found",
		})

		return
	}

	switch trackType {
	case "error":
		controllerError(c, &prj)
	case "log":
		controllerLog(c, &prj)
	default:
		c.JSON(common.Response{
			OK:    false,
			Error: "unknown type of track",
		})

		return
	}
}
