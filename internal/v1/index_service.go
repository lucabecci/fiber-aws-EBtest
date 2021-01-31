package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucabecci/fiber-aws-EBtest/pkg/home"
)

//IndexService is a one of my services of my api
type IndexService struct {
	Repository home.Repository
}

//Say is the first controller of this service
func (ir *IndexService) Say(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func (ir *IndexService) Service() *fiber.App {
	service := fiber.New()

	service.Get("/", ir.Say)

	return service
}
