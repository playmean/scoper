package database

import (
	"fmt"

	"github.com/playmean/scoper/config"
	"github.com/playmean/scoper/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// DBConn connection to database
	DBConn *gorm.DB

	tag = "DB"
)

// Init database and migrate models
func Init(conf *config.Config) error {
	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		conf.Database.Host,
		conf.Database.User,
		conf.Database.Password,
		conf.Database.DBName,
		conf.Database.Port,
	)

	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	logger.Log(tag, "connected")

	return nil
}
