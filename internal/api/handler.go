package api

import (
	"iBat/homework/pkg/tadaptor"
	"iBat/homework/pkg/validator"
	"iBat/homework/views/components"
	"time"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
)

type ApiHandler struct {
	router fiber.Router
}

func NewApiHandler(router fiber.Router) *ApiHandler {
	handler := &ApiHandler{
		router: router,
	}

	api := handler.router.Group("/api")
	api.Post("/register", handler.register)

	return handler
}

func (h *ApiHandler) register(c *fiber.Ctx) error {
	form := RegisterForm{
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
		comp = components.Notification("Пользователь успешно создан", components.NotificationSuccess)
	}
	time.Sleep(time.Second * 2)

	return tadaptor.Render(c, &comp)
}
