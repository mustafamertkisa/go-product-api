package router

import (
	"go-product-api/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, h *handler.ProductHandler) {
	api := app.Group("/api/products")
	api.Get("/", h.GetAll)
	api.Get("/:id", h.GetById)
	api.Post("/", h.Create)
	api.Put("/:id", h.Update)
	api.Delete("/:id", h.Delete)
}
