package main

import (
	"flag"

	"github.com/playmean/scoper/config"
	"github.com/playmean/scoper/connection"
	"github.com/playmean/scoper/database"
	"github.com/playmean/scoper/logger"
	"github.com/playmean/scoper/project"
	"github.com/playmean/scoper/router"
	"github.com/playmean/scoper/tag"
	"github.com/playmean/scoper/track"
	"github.com/playmean/scoper/user"

	"github.com/gofiber/fiber"
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

	app := fiber.New(&fiber.Settings{
		ServerHeader:          "scoper",
		StrictRouting:         true,
		DisableDefaultDate:    true,
		DisableStartupMessage: true,
	})

	router.Setup(conf, app)

	err = app.Listen(conf.Port)

	if err != nil {
		logger.Fatal("APP", "%v", err)
	}
}
