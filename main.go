package main

import (
	"flag"

	"git.playmean.xyz/playmean/error-tracking/config"
	"git.playmean.xyz/playmean/error-tracking/database"
	"git.playmean.xyz/playmean/error-tracking/logger"
	"git.playmean.xyz/playmean/error-tracking/project"
	"git.playmean.xyz/playmean/error-tracking/router"
	"git.playmean.xyz/playmean/error-tracking/track"
	"git.playmean.xyz/playmean/error-tracking/user"

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

	user.Populate(config.SuperUsers)

	app := fiber.New(&fiber.Settings{
		ServerHeader:          "error-tracking",
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
