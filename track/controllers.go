package track

import (
	"encoding/json"

	"git.playmean.xyz/playmean/error-tracking/database"
	"git.playmean.xyz/playmean/error-tracking/logger"
	"git.playmean.xyz/playmean/error-tracking/project"

	"github.com/gofiber/fiber"
)

func controllerError(c *fiber.Ctx, prj *project.Project) {
	db := database.DBConn

	var body reportPacket

	c.BodyParser(&body)

	formatted, _ := json.MarshalIndent(body, "", "    ")

	logger.Log("TRACK:ERROR", prj.Key, string(formatted))

	db.Create(&Track{
		Type:       "error",
		ProjectKey: prj.Key,

		Message:  body.Message,
		Stack:    body.Stack,
		Filename: body.Source.Filename,
		Lineno:   body.Source.Position.Lineno,
		Colno:    body.Source.Position.Colno,
	})

	c.JSON(response{
		OK: true,
	})
}

func controllerLog(c *fiber.Ctx, prj *project.Project) {
	db := database.DBConn

	var body interface{}

	c.BodyParser(&body)

	formatted, _ := json.MarshalIndent(body, "", "    ")
	formattedJSON := string(formatted)

	logger.Log("TRACK:LOG", prj.Key, formattedJSON)

	db.Create(&Track{
		Type:       "log",
		ProjectKey: prj.Key,

		Meta: formattedJSON,
	})

	c.JSON(response{
		OK: true,
	})
}
