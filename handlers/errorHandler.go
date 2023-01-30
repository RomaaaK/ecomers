package handlers

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	if code == 404 {
		c.SendStatus(fiber.StatusNotFound)
		return c.Render("404", fiber.Map{}, "")
	}

	c.SendString(fmt.Sprintf("Error: %v", code))

	return nil
}
