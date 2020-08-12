package user

import (
	"time"

	"git.playmean.xyz/playmean/error-tracking/database"
)

// User model
type User struct {
	ID           uint `gorm:"primary_key"`
	Username     string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Migrate user table
func Migrate() {
	database.DBConn.AutoMigrate(&User{})
}

// Authorizer for users
func Authorizer(superusers map[string]string) func(string, string) bool {
	return func(username, password string) bool {
		if _, ok := superusers[username]; ok {
			return superusers[username] == password
		}

		db := database.DBConn

		var found User

		db.First(&found, &User{
			Username:     username,
			PasswordHash: hashPassword(password),
		})

		return found.ID > 0
	}
}
