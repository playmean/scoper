package controllers

import (
	"encoding/json"

	"github.com/playmean/scoper/common"
	"github.com/playmean/scoper/database"
	"github.com/playmean/scoper/project"
	"github.com/playmean/scoper/track"
	"github.com/playmean/scoper/user"

	"github.com/gofiber/fiber"
)

// MiddlewareView method
func MiddlewareView(c *fiber.Ctx) {
	db := database.DBConn

	usr := c.Locals("user").(*user.User)

	projectKey := c.Params("key")

	var prj project.Project

	db.First(&prj, "key = ? AND (public = ? OR owner_id = ?)", projectKey, true, usr.ID)

	if prj.ID == 0 {
		c.Status(fiber.StatusNotFound).JSON(common.Response{
			OK:    false,
			Error: "project not found",
		})

		return
	}

	c.Locals("project", &prj)

	c.Next()
}

// ViewEnvironments method
func ViewEnvironments(c *fiber.Ctx) {
	db := database.DBConn

	prj := c.Locals("project").(*project.Project)

	list := make([]respEnvironment, 0)

	res := db.Raw(`
	SELECT
		t.environment as environment,
		COUNT(*) as count
	FROM tracks AS t
		INNER JOIN projects AS p ON (p.id = t.project_id)
	WHERE p.id = ?
	GROUP BY t.environment`, prj.ID).Scan(&list)

	common.Answer(c, res.Error, list)
}

// ViewTags method
func ViewTags(c *fiber.Ctx) {
	db := database.DBConn

	prj := c.Locals("project").(*project.Project)

	list := make([]respTagName, 0)

	res := db.Raw(`
	SELECT
		tn.id as id,
		tn.name as name,
		COUNT(DISTINCT tv.value) as count
	FROM tag_values AS tv
		INNER JOIN tag_names AS tn ON (tn.id = tv.name_id)
		INNER JOIN con_track_tags AS ctt ON (ctt.tag_value_id = tv.id)
		INNER JOIN tracks AS t ON (t.id = ctt.track_id)
		INNER JOIN projects AS p ON (p.id = t.project_id)
	WHERE p.id = ?
	GROUP BY tn.name`, prj.ID).Scan(&list)

	common.Answer(c, res.Error, list)
}

// ViewTagValues method
func ViewTagValues(c *fiber.Ctx) {
	db := database.DBConn

	prj := c.Locals("project").(*project.Project)

	tagNameID := c.Params("id")

	list := make([]respTagValue, 0)

	res := db.Raw(`
	SELECT
		tv.value as value,
		COUNT(*) as count
	FROM tag_values AS tv
		INNER JOIN tag_names AS tn ON (tn.id = tv.name_id)
		INNER JOIN con_track_tags AS ctt ON (ctt.tag_value_id = tv.id)
		INNER JOIN tracks AS t ON (t.id = ctt.track_id)
		INNER JOIN projects AS p ON (p.id = t.project_id)
	WHERE p.id = ? AND tn.id = ?
	GROUP BY tv.value`, prj.ID, tagNameID).Scan(&list)

	common.Answer(c, res.Error, list)
}

// ViewTracks method
func ViewTracks(c *fiber.Ctx) {
	db := database.DBConn

	prj := c.Locals("project").(*project.Project)

	var origTracks []track.Track
	var resTracks = make(map[string][]respTrack)

	res := db.Find(&origTracks, "archived = 0 AND project_id = ?", prj.ID)

	if res.Error != nil {
		common.Answer(c, res.Error, &origTracks)

		return
	}

	for _, trk := range origTracks {
		if _, ok := resTracks[trk.Type]; !ok {
			resTracks[trk.Type] = make([]respTrack, 0)
		}

		resTracks[trk.Type] = append(resTracks[trk.Type], respTrack{
			Environment: trk.Environment,

			Message:  trk.Message,
			Filename: trk.Filename,

			CreatedAt: trk.CreatedAt,
		})
	}

	common.Answer(c, res.Error, &resTracks)
}

// ViewTrack method
func ViewTrack(c *fiber.Ctx) {
	db := database.DBConn

	prj := c.Locals("project").(*project.Project)

	trackID := c.Params("id")

	var origTrack track.Track
	var resTrack respTrack

	res := db.First(&origTrack, "archived = 0 AND id = ? AND project_id = ?", trackID, prj.ID)

	if res.Error != nil {
		common.Answer(c, res.Error, &origTrack)

		return
	}

	resTrack.Type = origTrack.Type
	resTrack.Environment = origTrack.Environment

	resTrack.Message = origTrack.Message
	resTrack.Stack = origTrack.Stack
	resTrack.Filename = origTrack.Filename
	resTrack.Lineno = origTrack.Lineno
	resTrack.Colno = origTrack.Colno

	resTrack.CreatedAt = origTrack.CreatedAt

	json.Unmarshal([]byte(origTrack.Meta), &resTrack.Meta)

	res = db.Raw(`
	SELECT
		tn.name as name,
		tv.value as value
	FROM tag_values AS tv
		INNER JOIN tag_names AS tn ON (tn.id = tv.name_id)
		INNER JOIN con_track_tags AS ctt ON (ctt.tag_value_id = tv.id)
		INNER JOIN tracks AS t ON (t.id = ctt.track_id)
		INNER JOIN projects AS p ON (p.id = t.project_id)
	WHERE p.id = ? AND t.id = ?`, prj.ID, trackID).Scan(&resTrack.Tags)

	common.Answer(c, res.Error, &resTrack)
}
