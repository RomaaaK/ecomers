package handlers

import (
	"example/ecomers/models"
	"example/ecomers/services"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	var brands []models.Brand

	services.DB.Find(&brands)

	return c.Render("index", fiber.Map{
		"Slider": true,
		"Brands": brands,
	})
}
