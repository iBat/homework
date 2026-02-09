package api

import (
	"iBat/homework/internal/users"
	"iBat/homework/pkg/tadaptor"
	"iBat/homework/pkg/validator"
	"iBat/homework/views/components"
	"log/slog"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type ApiHandler struct {
	router   fiber.Router
	userRepo *users.UserRepository
	logger   *slog.Logger
	session  *session.Store
}

func NewApiHandler(router fiber.Router, userRepo *users.UserRepository, logger *slog.Logger, session *session.Store) *ApiHandler {
	handler := &ApiHandler{
		router:   router,
		userRepo: userRepo,
		logger:   logger,
		session:  session,
	}

	api := handler.router.Group("/api")
	api.Post("/register", handler.register)
	api.Post("/login", handler.login)
	api.Get("/logout", handler.logout)

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

func (h *ApiHandler) login(c *fiber.Ctx) error {
	form := LoginForm{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Не задан или неверен"},
		&validators.StringIsPresent{Name: "Password", Field: form.Password},
	)

	var comp templ.Component
	if len(errors.Errors) > 0 {
		comp = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
	} else {
		user, err := h.userRepo.ValidateUserPassword(form.Email, form.Password)
		if err != nil {
			comp = components.Notification("Ошибка при входе пользователя", components.NotificationFail)
		} else {
			sess, err := h.session.Get(c)
			if err != nil {
				panic(err)
			}
			sess.Set("email", form.Email)
			if err := sess.Save(); err != nil {
				panic(err)
			}
			h.logger.Info("User logged in successfully", slog.String("email", user.Email))
			c.Response().Header.Set("HX-Redirect", "/")
			return c.Redirect("/", http.StatusOK)
		}
	}

	return tadaptor.Render(c, &comp)
}

func (h *ApiHandler) logout(c *fiber.Ctx) error {
	sess, err := h.session.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Destroy()
	c.Response().Header.Set("HX-Redirect", "/")
	return c.Redirect("/", http.StatusOK)
}
