package user

import (
	"time"

	"git.playmean.xyz/playmean/error-tracking/database"
)

// User model
type User struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

// Migrate table
func Migrate() {
	db := database.DBConn

	if db.HasTable(&User{}) {
		return
	}

	db.CreateTable(&User{})
}

// Populate superusers
func Populate(superusers map[string]string) {
	db := database.DBConn

	for username, password := range superusers {
		var found User

		db.FirstOrCreate(&found, &User{
			Username:     username,
			PasswordHash: hashPassword(password),
			Role:         "super",
		})
	}
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
