package controllers

import (
	"git.playmean.xyz/playmean/scoper/connection"
	"git.playmean.xyz/playmean/scoper/database"
	"git.playmean.xyz/playmean/scoper/tag"
	"git.playmean.xyz/playmean/scoper/track"
)

func newTags(track *track.Track, tags map[string]string) []uint {
	db := database.DBConn

	createdList := make([]uint, 0)

	for name, value := range tags {
		tagName := tag.Name{Name: name}

		db.FirstOrCreate(&tagName, &tagName)

		tagValue := tag.Value{Name: tagName, Value: value}

		db.FirstOrCreate(&tagValue, &tagValue)

		conn := connection.ConTrackTag{
			Track:    *track,
			TagValue: tagValue,
		}

		db.Save(&conn)

		createdList = append(createdList, tagValue.ID)
	}

	return createdList
}
