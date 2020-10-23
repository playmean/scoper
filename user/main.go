package user

import (
	"crypto/sha256"
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

		res := db.First(&found, &User{
			Username: username,
		})

		if res.Error != nil {
			db.Create(&User{
				Username:     username,
				PasswordHash: HashPassword(password),
				Role:         "super",
			})
		}
	}
}

// HashPassword string
func HashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))

	return hex.EncodeToString(hasher.Sum(nil))
}
