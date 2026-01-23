package main

import (
	"iBat/homework/config"
	"iBat/homework/internal/api"
	"iBat/homework/internal/pages"
	"iBat/homework/internal/users"
	"iBat/homework/pkg/database"
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
	dbConfig := config.NewDatabaseConfig()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app.Use(slogfiber.New(logger))
	app.Use(recover.New())
	app.Static("/public", "./public")
	dbpool := database.CreateDbPool(dbConfig, logger)
	defer dbpool.Close()

	// Initialize repositories
	userRepository := users.NewUserRepository(dbpool, logger)

	// Initialize handlers
	pages.NewPagesHandler(app)
	api.NewApiHandler(app, userRepository, logger)
	users.NewUserHandler(app, userRepository, logger)

	if err := app.Listen(":3000"); err != nil {
		slog.Error("Failed to start server", slog.String("error", err.Error()))
	}
}
