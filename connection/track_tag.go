package connection

import (
	"git.playmean.xyz/playmean/scoper/tag"
	"git.playmean.xyz/playmean/scoper/track"
)

// ConTrackTag connection between two objects
type ConTrackTag struct {
	ID uint `json:"id" gorm:"primary_key"`

	TrackID uint
	Track   track.Track

	TagValueID uint
	TagValue   tag.Value

	CreatedAt int64 `json:"-"`
	UpdatedAt int64 `json:"-"`
}
