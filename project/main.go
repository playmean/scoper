package project

import (
	"time"

	"git.playmean.xyz/playmean/error-tracking/database"
	"git.playmean.xyz/playmean/error-tracking/user"
)

// Project model
type Project struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Owner     user.User `json:"owner"`
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
