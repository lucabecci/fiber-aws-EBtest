package internal

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

//Server is a struct for the server app
type Server struct {
	srv *fiber.App
}

//New is for the start the srv
func New(port string) (*Server, error) {
	srv := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Test",
		ReadTimeout:   10 * time.Second,
		WriteTimeout:  10 * time.Second,
	})

	return &Server{srv}, nil
}
