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
	var form RegisterForm
	if err := c.BodyParser(&form); err != nil {
		h.logger.Error("Failed to parse registration form", slog.String("error", err.Error()))
		return c.Status(fiber.StatusBadRequest).SendString("Invalid form data")
	}

	email, err := h.repository.CreateUser(form)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to register user")
	}

	return c.SendString(email)
}
