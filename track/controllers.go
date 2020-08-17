package track

import (
	"encoding/json"

	"git.playmean.xyz/playmean/error-tracking/common"
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

	track := Track{
		Type:        "error",
		ProjectKey:  prj.Key,
		Environment: resolveEnvironment(c),

		Message:  body.Message,
		Stack:    body.Stack,
		Filename: body.Source.Filename,
		Lineno:   body.Source.Position.Lineno,
		Colno:    body.Source.Position.Colno,

		Tags: marshal(body.Tags),
	}

	db.Save(&track)

	c.JSON(common.Response{
		OK: true,
		Data: map[string]string{
			"hash": hashID(track.ID),
		},
	})
}

func controllerLog(c *fiber.Ctx, prj *project.Project) {
	db := database.DBConn

	var body logPacket

	c.BodyParser(&body)

	formatted, _ := json.MarshalIndent(body, "", "    ")

	logger.Log("TRACK:LOG", prj.Key, string(formatted))

	track := Track{
		Type:        "log",
		ProjectKey:  prj.Key,
		Environment: resolveEnvironment(c),

		Meta: marshal(body.Data),
		Tags: marshal(body.Tags),
	}

	db.Save(&track)

	c.JSON(common.Response{
		OK: true,
		Data: map[string]string{
			"hash": hashID(track.ID),
		},
	})
}
