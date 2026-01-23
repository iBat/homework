package api

import (
	"iBat/homework/internal/users"
	"iBat/homework/pkg/tadaptor"
	"iBat/homework/pkg/validator"
	"iBat/homework/views/components"
	"log/slog"
	"time"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
)

type ApiHandler struct {
	router   fiber.Router
	userRepo *users.UserRepository
	logger   *slog.Logger
}

func NewApiHandler(router fiber.Router, userRepo *users.UserRepository, logger *slog.Logger) *ApiHandler {
	handler := &ApiHandler{
		router:   router,
		userRepo: userRepo,
		logger:   logger,
	}

	api := handler.router.Group("/api")
	api.Post("/register", handler.register)

	return handler
}

func (h *ApiHandler) register(c *fiber.Ctx) error {
	form := users.RegisterForm{
		Name:     c.FormValue("name"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	errors := validate.Validate(
		&validators.StringIsPresent{Name: "Name", Field: form.Name},
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Не задан или неверен"},
		&validators.StringIsPresent{Name: "Password", Field: form.Password},
	)

	var comp templ.Component
	if len(errors.Errors) > 0 {
		comp = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
	} else {
		_, err := h.userRepo.CreateUser(form)
		if err != nil {
			comp = components.Notification("Ошибка при создании пользователя", components.NotificationFail)
		} else {
			comp = components.Notification("Пользователь успешно создан", components.NotificationSuccess)
		}
	}
	time.Sleep(time.Second * 2)

	return tadaptor.Render(c, &comp)
}
