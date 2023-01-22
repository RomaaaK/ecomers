package handlers

import (
	"example/ecomers/models"
	"example/ecomers/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Product(c *fiber.Ctx) error {
	var brands []models.Brand
	var categories []models.Category
	var product models.Product

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(404)
	}

	services.DB.Model(&models.Category{}).Preload("Childrens").Find(&categories)
	services.DB.Find(&brands)
	services.DB.First(&product, id)

	if product.ID == 0 {
		return c.SendStatus(404)
	}

	var productBrand models.Brand

	for _, b := range brands {
		if b.ID == *product.BrandID {
			productBrand = b
			break
		}
	}

	return c.Render("product-details", fiber.Map{
		"Brands":       brands,
		"Categories":   categories,
		"Product":      product,
		"ProductBrand": productBrand,
	})
}
