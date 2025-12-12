package main

import (
	"iBat/homework/config"
	"iBat/homework/internal/pages"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	app := fiber.New()

	app.Use(recover.New())
	pages.NewPagesHandler(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
