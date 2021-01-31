package main

import (
	"log"
	"os"
	"os/signal"

	_ "github.com/joho/godotenv/autoload"
	"github.com/lucabecci/fiber-aws-EBtest/internal"
)

func main() {
	server, err := internal.New()

	if err != nil {
		log.Panic(err.Error())
	}
	port := os.Getenv("port")
	server.Start(port)

	//If the developer use the Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	server.Close()
}
