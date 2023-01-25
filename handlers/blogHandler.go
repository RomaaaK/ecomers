package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Blog(c *fiber.Ctx) error {

	return c.Render("blog/index", fiber.Map{})
}

func BlogById(c *fiber.Ctx) error {

	return c.Render("blog/blog-single", fiber.Map{})
}
