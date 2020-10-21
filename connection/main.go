package connection

import "git.playmean.xyz/playmean/scoper/database"

// Migrate tables
func Migrate() {
	db := database.DBConn

	db.AutoMigrate(&ConTrackTag{})
}
