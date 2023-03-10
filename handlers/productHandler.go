package handlers

import (
	"example/ecomers/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Product(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound)
	}

	product, err := services.GetProductById(id)

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound)
	}

	return c.Render("product-details", fiber.Map{
		"Product": product,
	})
}

func ProductByCategory(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound)
	}

	return c.Render("index", fiber.Map{
		"Products": services.GetProductsByCategory(id),
	})
}

func ProductByBrand(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound)
	}

	return c.Render("index", fiber.Map{
		"Products": services.GetProductsByBrand(id),
	})
}
