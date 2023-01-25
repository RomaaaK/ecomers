package handlers

import (
	"example/ecomers/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Blog(c *fiber.Ctx) error {

	return c.Render("blog/index", fiber.Map{
		"Posts": services.GetBlogs(),
	})
}

func BlogById(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Next()
	}

	post, err := services.GetBlogById(id)

	if err != nil {
		return c.Next()
	}

	return c.Render("blog/blog-single", fiber.Map{
		"Post": post,
	})
}
