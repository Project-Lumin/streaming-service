package main

import (
	"fmt"
	"log"
	"streaming-service/config"
	"streaming-service/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	fmt.Println("starting server")

	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024, // 100 MB
	})
	app.Use(cors.New(cors.Config{
    AllowOrigins: "*",
    AllowMethods: "GET,HEAD,OPTIONS",
    AllowHeaders: "*",
}))

	app.Use(logger.New())
	config.Load()
	config.ConnectDatabase()
	router.Route(app)

	log.Fatal(app.Listen(":8000"))
}
