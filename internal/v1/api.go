package v1

import (
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	index := fiber.New()
	index.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	return index
}
