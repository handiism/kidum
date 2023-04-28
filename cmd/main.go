package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/handiism/kidum/api/route"
	"github.com/handiism/kidum/bootstrap"
)

func main() {
	app := bootstrap.App()
	defer app.Close()

	fib := fiber.New(fiber.Config{
		BodyLimit: 5 * 1024 * 1024,
	})

	route.Setup(fib, app.PgPool)

	if err := fib.Listen(app.Env.ServerAddress); err != nil {
		log.Fatal(err.Error())
	}
}
