package handlers

import (
	"example/ecomers/helpers"
	"example/ecomers/models"
	"example/ecomers/services"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	var products []models.Product

	services.DB.Find(&products)

	return c.Render("index", helpers.PreloadMainLayoutData(fiber.Map{
		"Slider":   true,
		"Products": products,
	}))
}
