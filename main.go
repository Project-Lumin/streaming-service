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

	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024, // 100 MB
	})
	config.Load()
	config.ConnectDatabase()
	router.Route(app)

	log.Fatal(app.Listen(":8000"))
}
