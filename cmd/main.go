package main

import (
	"iBat/homework/config"
	"iBat/homework/internal/pages"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	config.Init()
	engine := html.New("./html", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app.Use(slogfiber.New(logger))
	app.Use(recover.New())
	app.Static("/public", "./public")

	pages.NewPagesHandler(app)

	if err := app.Listen(":3000"); err != nil {
		slog.Error("Failed to start server", slog.String("error", err.Error()))
	}
}
