package connection

import "github.com/playmean/scoper/database"

// Migrate tables
func Migrate() {
	db := database.DBConn

	db.AutoMigrate(&ConTrackTag{})
}
