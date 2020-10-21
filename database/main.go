package database

import (
	"git.playmean.xyz/playmean/scoper/config"
	"git.playmean.xyz/playmean/scoper/logger"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var (
	// DBConn connection to database
	DBConn *gorm.DB

	tag = "DB"
)

// Init database and migrate models
func Init(conf *config.Config) error {
	var err error

	DBConn, err = gorm.Open(sqlite.Open(conf.Database), &gorm.Config{})

	if err != nil {
		return err
	}

	logger.Log(tag, "connected")

	return nil
}
