package user

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/playmean/scoper/database"
)

// User model
type User struct {
	ID uint `json:"id" gorm:"primary_key"`

	Username     string `json:"username"`
	FullName     string `json:"fullname"`
	Password     string `json:"password,omitempty" gorm:"-"`
	PasswordHash string `json:"-"`
	Role         string `json:"role"`

	CreatedAt int64 `json:"-"`
	UpdatedAt int64 `json:"-"`
}

// Migrate table
func Migrate() {
	db := database.DBConn

	db.AutoMigrate(&User{})
}

// Populate superusers
func Populate(superusers map[string]string) {
	db := database.DBConn

	for username, password := range superusers {
		var found User

		db.FirstOrCreate(&found, &User{
			Username:     username,
			PasswordHash: HashPassword(password),
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
			PasswordHash: HashPassword(password),
		})

		return found.ID > 0
	}
}

// HashPassword string
func HashPassword(password string) string {
	hasher := sha1.New()
	hasher.Write([]byte(password))

	return hex.EncodeToString(hasher.Sum(nil))
}
