package main

import (
	"example/ecomers/handlers"
	"example/ecomers/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	mainLayout := app.Group("/")

	mainLayout.Use(middleware.BindMainLayoutData)

	mainLayout.Get("/", handlers.Index)
	mainLayout.Get("/product/:id<int>", handlers.Product)
	mainLayout.Get("/category/:id<int>", handlers.ProductByCategory)
	mainLayout.Get("/brand/:id<int>", handlers.ProductByBrand)

	blog := mainLayout.Group("/blog")
	blog.Get("/", handlers.Blog)
	blog.Get("/:id<int>", handlers.BlogById)
}
