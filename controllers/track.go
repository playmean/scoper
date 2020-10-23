package controllers

import (
	"encoding/json"

	"github.com/playmean/scoper/common"
	"github.com/playmean/scoper/database"
	"github.com/playmean/scoper/logger"
	"github.com/playmean/scoper/project"
	"github.com/playmean/scoper/track"

	"github.com/gofiber/fiber"
)

// MiddlewareTrack for tracking APIs
func MiddlewareTrack(c *fiber.Ctx) {
	db := database.DBConn

	var prj project.Project

	projectKey := c.Params("key")
	trackType := c.Params("type")

	db.First(&prj, "archived = 0 AND key = ?", projectKey)

	if prj.ID == 0 {
		c.Status(fiber.StatusNotFound).JSON(common.Response{
			OK:    false,
			Error: "project not found",
		})

		return
	}

	switch trackType {
	case "error":
		trackError(c, &prj)
	case "log":
		trackLog(c, &prj)
	default:
		c.Status(fiber.StatusBadRequest).JSON(common.Response{
			OK:    false,
			Error: "unknown type of track",
		})

		return
	}
}

func trackError(c *fiber.Ctx, prj *project.Project) {
	db := database.DBConn

	var body track.ReportPacket

	c.BodyParser(&body)

	formatted, _ := json.MarshalIndent(body, "", "    ")

	logger.Log("TRACK:ERROR", prj.Key, string(formatted))

	track := &track.Track{
		Type:        "error",
		Project:     *prj,
		Environment: track.ResolveEnvironment(c),

		Message:  body.Message,
		Stack:    body.Stack,
		Filename: body.Source.Filename,
		Lineno:   body.Source.Position.Lineno,
		Colno:    body.Source.Position.Colno,
	}

	res := db.Save(track)

	newTags(track, body.Tags)

	common.Answer(c, res.Error, map[string]string{
		"hash": hashID(track.ID),
	})
}

func trackLog(c *fiber.Ctx, prj *project.Project) {
	db := database.DBConn

	var body track.LogPacket

	c.BodyParser(&body)

	formatted, _ := json.MarshalIndent(body, "", "    ")

	logger.Log("TRACK:LOG", prj.Key, string(formatted))

	track := &track.Track{
		Type:        "log",
		Project:     *prj,
		Environment: track.ResolveEnvironment(c),

		Meta: marshal(body.Data),
	}

	res := db.Save(track)

	newTags(track, body.Tags)

	common.Answer(c, res.Error, map[string]string{
		"hash": hashID(track.ID),
	})
}
