package project

import (
	"github.com/playmean/scoper/database"
)

// Project model
type Project struct {
	ID uint `json:"id" gorm:"primary_key"`

	Key     string `json:"key"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	OwnerID uint   `json:"owner_id"`
	Public  bool   `json:"public"`

	Archived  int64 `json:"-"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"-"`
}

// Migrate table
func Migrate() {
	db := database.DBConn

	db.AutoMigrate(&Project{})
}
