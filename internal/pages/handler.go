package pages

import (
	"iBat/homework/pkg/tadaptor"
	"iBat/homework/views"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type PagesHandler struct {
	router     fiber.Router
	categories []string
	session    *session.Store
}

func NewPagesHandler(router fiber.Router, session *session.Store) *PagesHandler {
	handler := &PagesHandler{
		router:  router,
		session: session,
		categories: []string{
			"Еда",
			"Животные",
			"Машины",
			"Спорт",
			"Музыка",
			"Технологии",
			"Прочее",
		},
	}

	handler.router.Get("/", handler.HomePage)
	handler.router.Get("/register", handler.Register)
	handler.router.Get("/login", handler.Login)

	return handler
}

func (h *PagesHandler) HomePage(c *fiber.Ctx) error {
	component := views.Main(h.categories)

	return tadaptor.Render(c, &component)
}

func (h *PagesHandler) Register(c *fiber.Ctx) error {
	component := views.Register()

	return tadaptor.Render(c, &component)
}

func (h *PagesHandler) Login(c *fiber.Ctx) error {
	component := views.Login()

	return tadaptor.Render(c, &component)
}
