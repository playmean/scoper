package track

import (
	"github.com/playmean/scoper/database"
	"github.com/playmean/scoper/project"
)

// Track model
type Track struct {
	ID uint `json:"id" gorm:"primary_key"`

	Type        string `json:"type"`
	ProjectID   uint
	Project     project.Project
	Environment string `json:"environment"`

	Message  string `json:"message,omitempty" gorm:"type:text"`
	Stack    string `json:"stack,omitempty" gorm:"type:text"`
	Filename string `json:"filename,omitempty"`
	Lineno   int    `json:"lineno,omitempty"`
	Colno    int    `json:"colno,omitempty"`

	Meta string `json:"meta,omitempty" gorm:"type:json"`

	Archived  int64 `json:"-"`
	CreatedAt int64 `json:"-"`
	UpdatedAt int64 `json:"-"`
}

// Migrate table
func Migrate() {
	db := database.DBConn

	db.AutoMigrate(&Track{})
}
