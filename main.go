package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"streaming-service/config"
	"streaming-service/router"
)

func main() {
	fmt.Println("starting server")

	app := fiber.New()
	config.Load()
	router.Route(app)

	log.Fatal(app.Listen(":8000"))
}
