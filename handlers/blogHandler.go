package handlers

import (
	"example/ecomers/models"
	"example/ecomers/services"

	"github.com/gofiber/fiber/v2"
)

func Blog(c *fiber.Ctx) error {

	var categories []models.Category
	var brands []models.Brand

	services.DB.Model(&models.Category{}).Preload("Childrens").Find(&categories)
	services.DB.Find(&brands)

	return c.Render("blog/index", fiber.Map{
		"Brands":     brands,
		"Categories": categories,
	})
}

func BlogById(c *fiber.Ctx) error {

	var categories []models.Category
	var brands []models.Brand

	services.DB.Model(&models.Category{}).Preload("Childrens").Find(&categories)
	services.DB.Find(&brands)

	return c.Render("blog/blog-single", fiber.Map{
		"Brands":     brands,
		"Categories": categories,
	})
}
