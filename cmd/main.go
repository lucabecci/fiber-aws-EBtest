package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/lucabecci/fiber-aws-EBtest/internal"
)

func main() {
	fmt.Println("Hello world")
	server, err := internal.New()

	if err != nil {
		log.Panic(err.Error())
	}

	server.Start()

	//If the developer use the Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	server.Close()
}
