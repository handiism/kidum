package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/handiism/kidum/bootstrap"
)

func main() {
	app := bootstrap.App()
	defer app.Close()

	fiber := fiber.New()

	if err := fiber.Listen(app.Env.ServerAddress); err != nil {
		log.Fatal(err.Error())
	}
}
