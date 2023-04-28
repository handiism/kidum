package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/handiism/kidum/api/controller"
	"github.com/handiism/kidum/api/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Setup(app *fiber.App, pool *pgxpool.Pool) {
	controller := controller.NewController(pool)
	middleware := middleware.NewMiddleware()

	forms := app.Group("/forms")
	forms.Post("/", middleware.Multipart.Basic(), controller.Form.Insert())
}
