package project

import (
	"time"

	"git.playmean.xyz/playmean/error-tracking/database"
)

// Project model
type Project struct {
	ID uint `json:"id" gorm:"primary_key"`

	Key     string `json:"key"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	OwnerID uint   `json:"owner_id"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Migrate table
func Migrate() {
	db := database.DBConn

	if db.HasTable(&Project{}) {
		return
	}

	db.CreateTable(&Project{})
}
