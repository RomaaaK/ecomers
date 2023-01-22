package main

import (
	"example/ecomers/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", handlers.Index)
	app.Get("/product/:id<int>", handlers.Product)
}
