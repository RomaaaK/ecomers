package handlers

import (
	"example/ecomers/models"
	"example/ecomers/services"

	"github.com/gofiber/fiber/v2"
)

func Product(c *fiber.Ctx) error {
	var brands []models.Brand

	services.DB.Find(&brands)
	return c.Render("product-details", fiber.Map{
		"Brands": brands,
	})
}
