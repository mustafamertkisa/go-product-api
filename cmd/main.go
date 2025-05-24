package main

import (
	"go-product-api/config"
	"go-product-api/handler"
	"go-product-api/repository"
	"go-product-api/router"
	"go-product-api/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.Connect()

	repo := repository.NewProductRepository(config.ProductCollection)
	svc := service.NewProductService(repo)
	handler := handler.NewProductHandler(svc)

	router.SetupRoutes(app, handler)

	app.Listen(":3000")
}
