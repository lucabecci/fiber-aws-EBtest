package home

import "github.com/gofiber/fiber/v2"

//Repository is a repositroy of the index api
type Repository interface {
	Say(c *fiber.Ctx) error
}
