package internal

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	v1 "github.com/lucabecci/fiber-aws-EBtest/internal/v1"
)

//Server is a struct for the server app
type Server struct {
	server *fiber.App
}
type Versions struct {
	V1 string
}

//New is for the start the srv
func New() (*Server, error) {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Test",
		ReadTimeout:   10 * time.Second,
		WriteTimeout:  10 * time.Second,
	})
	//routes
	v1 := v1.New()
	app.Mount("/v1", v1.Router)
	app.Get("/", Index)
	server := Server{server: app}
	return &server, nil
}

func (serv *Server) Close() error {
	log.Printf("Server stopped")
	return nil
}

func (serv *Server) Start(port string) {
	log.Printf("Server running on http://localhost:" + port)
	log.Fatal(serv.server.Listen(":" + port))
}

func Index(c *fiber.Ctx) error {
	var versions = Versions{
		V1: "http://localhost:4000/v1",
	}

	result, err := json.Marshal(versions)
	if err != nil {
		log.Panic(err.Error())
	}
	return c.Send(result)
}
