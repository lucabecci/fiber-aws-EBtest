package v1

import (
	"github.com/gofiber/fiber/v2"
)

type IndexRouter struct {
	Router *fiber.App
}

func New() *IndexRouter {
	index := fiber.New()

	is := &IndexService{}
	isService := is.Service()

	index.Mount("/", isService)

	return &IndexRouter{Router: index}
}
