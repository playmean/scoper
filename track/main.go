package track

import (
	"time"

	"git.playmean.xyz/playmean/error-tracking/database"
)

// Track model
type Track struct {
	ID uint `json:"id" gorm:"primary_key"`

	Type        string `json:"type"`
	ProjectKey  string `json:"project_key"`
	Environment string `json:"environment"`

	Message  string `json:"message,omitempty" gorm:"type:text"`
	Stack    string `json:"stack,omitempty" gorm:"type:text"`
	Filename string `json:"filename,omitempty"`
	Lineno   int    `json:"lineno,omitempty"`
	Colno    int    `json:"colno,omitempty"`

	Meta string `json:"meta,omitempty" gorm:"type:json"`
	Tags string `json:"tags,omitempty" gorm:"type:json"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Migrate table
func Migrate() {
	db := database.DBConn

	if db.HasTable(&Track{}) {
		return
	}

	db.CreateTable(&Track{})
}
