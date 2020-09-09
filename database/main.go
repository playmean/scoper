package database

import (
	"git.playmean.xyz/playmean/scoper/config"
	"git.playmean.xyz/playmean/scoper/logger"

	"github.com/jinzhu/gorm"

	// sqlite driver for gorm
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	// DBConn connection to database
	DBConn *gorm.DB

	tag = "DB"
)

// Init database and migrate models
func Init(conf *config.Config) error {
	var err error

	DBConn, err = gorm.Open("sqlite3", conf.Database)

	if err != nil {
		return err
	}

	logger.Log(tag, "connected")

	return nil
}
