package tag

import (
	"git.playmean.xyz/playmean/scoper/database"
)

// Name model
type Name struct {
	ID uint `json:"id" gorm:"primary_key"`

	Name string `json:"name"`

	CreatedAt int64 `json:"-"`
	UpdatedAt int64 `json:"-"`
}

// Value model
type Value struct {
	ID uint `json:"id,omitempty" gorm:"primary_key"`

	NameID uint
	Name   Name
	Value  string `json:"value"`

	CreatedAt int64 `json:"-"`
	UpdatedAt int64 `json:"-"`
}

// TableName for tag names
func (Name) TableName() string {
	return "tag_names"
}

// TableName for tag values
func (Value) TableName() string {
	return "tag_values"
}

// Migrate tables
func Migrate() {
	db := database.DBConn

	db.AutoMigrate(&Name{})
	db.AutoMigrate(&Value{})
}
