package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err.Error())
	}
}
