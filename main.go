package main

import (
	"fmt"
	"log"
	router "streaming-service/Router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("starting server")

	app := fiber.New()

	router.Route(app)

	log.Fatal(app.Listen(":8080"))
}
