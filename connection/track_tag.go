package connection

import (
	"github.com/playmean/scoper/tag"
	"github.com/playmean/scoper/track"
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
