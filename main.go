package main

import (
	"encoding/json"
	"flag"
	"log"

	"git.playmean.xyz/playmean/error-tracking/config"

	"github.com/gofiber/basicauth"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/helmet"
)

type reportPacketSourcePosition struct {
	Lineno int `json:"lineno"`
	Colno  int `json:"colno"`
}

type reportPacketSource struct {
	Filename string                     `json:"filename"`
	Position reportPacketSourcePosition `json:"position"`
}

type reportPacket struct {
	Message string             `json:"message"`
	Stack   string             `json:"stack"`
	Source  reportPacketSource `json:"source"`
}

// Response from REST API
type Response struct {
	OK    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

func main() {
	configPath := flag.String("config", "config.json", "app configuration file")

	flag.Parse()

	conf, err := config.Load(*configPath)

	if err != nil {
		log.Fatalln("[CONFIG]", err)
	}

	conf.Dump()

	app := fiber.New(&fiber.Settings{
		ServerHeader:          "error-tracking",
		StrictRouting:         true,
		DisableDefaultDate:    true,
		DisableStartupMessage: true,
	})

	users := make(map[string]string)

	for username, user := range conf.Users {
		users[username] = user.Password
	}

	app.Use(middleware.Compress())
	app.Use(middleware.Favicon())
	app.Use(helmet.New())

	app.Group("/api", basicauth.New(basicauth.Config{
		Users: users,
	}))

	app.Post("/track/:key", func(c *fiber.Ctx) {
		projectKey := c.Params("key")

		if projectKey == "" {
			c.JSON(Response{
				OK:    false,
				Error: "project key not specified",
			})

			return
		}

		var body reportPacket

		c.BodyParser(&body)

		formatted, _ := json.MarshalIndent(body, "", "    ")

		log.Println(projectKey, string(formatted))

		c.JSON(Response{
			OK: true,
		})
	})

	app.Listen(conf.Port)
}
