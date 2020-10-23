package user

import (
	"github.com/playmean/scoper/config"
	"github.com/playmean/scoper/database"
)

// Authorize user
func Authorize(username, password string) bool {
	if _, ok := config.SuperUsers[username]; ok {
		return config.SuperUsers[username] == password
	}

	db := database.DBConn

	var found User

	db.First(&found, &User{
		Username:     username,
		PasswordHash: HashPassword(password),
	})

	return found.ID > 0
}
