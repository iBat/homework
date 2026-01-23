package users

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	router     fiber.Router
	repository *UserRepository
	logger     *slog.Logger
}

func NewUserHandler(router fiber.Router, repository *UserRepository, logger *slog.Logger) *UserHandler {
	handler := &UserHandler{
		router:     router,
		repository: repository,
		logger:     logger,
	}
	handler.registerRoutes()
	return handler
}

func (h *UserHandler) registerRoutes() {
	h.router.Post("/register", h.handleRegister)
}

func (h *UserHandler) handleRegister(c *fiber.Ctx) error {
	form := RegisterForm{
		Name:     c.FormValue("name"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	email, err := h.repository.CreateUser(form)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to register user")
	}

	return c.SendString(email)
}
