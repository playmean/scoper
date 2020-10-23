package main

import (
	"flag"
	"strconv"

	"github.com/playmean/scoper/config"
	"github.com/playmean/scoper/connection"
	"github.com/playmean/scoper/database"
	"github.com/playmean/scoper/logger"
	"github.com/playmean/scoper/project"
	"github.com/playmean/scoper/router"
	"github.com/playmean/scoper/tag"
	"github.com/playmean/scoper/track"
	"github.com/playmean/scoper/user"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configPath := flag.String("config", "config.json", "app configuration file")

	flag.Parse()

	conf, err := config.Load(*configPath)

	if err != nil {
		logger.Fatal("CONFIG", "%v", err)
	}

	conf.Dump()

	err = database.Init(conf)

	if err != nil {
		logger.Fatal("DB", "%v", err)
	}

	user.Migrate()
	project.Migrate()
	track.Migrate()
	tag.Migrate()
	connection.Migrate()

	user.Populate(config.SuperUsers)

	app := fiber.New(fiber.Config{
		ServerHeader:          "scoper",
		StrictRouting:         true,
		DisableDefaultDate:    true,
		DisableStartupMessage: true,
	})

	router.Setup(conf, app)

	err = app.Listen(conf.Address + ":" + strconv.Itoa(conf.Port))

	if err != nil {
		logger.Fatal("APP", "%v", err)
	}
}
