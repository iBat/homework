package pages

import "github.com/gofiber/fiber/v2"

type PagesHandler struct {
	router fiber.Router
}

func NewPagesHandler(router fiber.Router) *PagesHandler {
	handler := &PagesHandler{
		router: router,
	}

	handler.router.Get("/", handler.HomePage)

	return handler
}

func (h *PagesHandler) HomePage(c *fiber.Ctx) error {
	return c.SendString("Welcome to the Home Page!")
}
