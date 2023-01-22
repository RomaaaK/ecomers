package handlers

import (
	"example/ecomers/models"
	"example/ecomers/services"

	"github.com/gofiber/fiber/v2"
)

func Product(c *fiber.Ctx) error {
	var brands []models.Brand
	var categories []models.Category

	services.DB.Model(&models.Category{}).Preload("Childrens").Find(&categories)
	services.DB.Find(&brands)
	return c.Render("product-details", fiber.Map{
		"Brands":     brands,
		"Categories": categories,
	})
}
