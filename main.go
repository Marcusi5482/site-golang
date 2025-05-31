package main

import (
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/:name?", func(c fiber.Ctx) error {
		return c.SendString("Hello World!\nGigachad\nhttps://github.com/Marcusi5482/site-golang/releases/tag/1")
	})

	app.Listen(":10000")
}
