package pages

import (
	"iBat/homework/pkg/tadaptor"
	"iBat/homework/views"

	"github.com/gofiber/fiber/v2"
)

type PagesHandler struct {
	router     fiber.Router
	categories []string
}

func NewPagesHandler(router fiber.Router) *PagesHandler {
	handler := &PagesHandler{
		router: router,
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

	return handler
}

func (h *PagesHandler) HomePage(c *fiber.Ctx) error {
	component := views.Main(h.categories)

	return tadaptor.Render(c, &component)
}
