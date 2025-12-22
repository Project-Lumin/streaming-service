package router

import (
	"streaming-service/api"

	"github.com/gofiber/fiber/v2"
)

func Route(g *fiber.App) {
	v1 := g.Group("/v1")
	v1.Get("/fetch/segment", api.FetchSegmentHandler)
}
