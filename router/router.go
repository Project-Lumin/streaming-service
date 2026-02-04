package router

import (
	"streaming-service/api"
	"streaming-service/config"

	"github.com/gofiber/fiber/v2"
)

func Route(g *fiber.App) {
	v1 := g.Group("/v1")

	v1.Get("/fetch/:videoId/*", api.FetchSegmentHandler)
	v1.Post("/upload", api.UploadVideoHandler)
	v1.Post("/video", api.CreateVideoDataHandler)
	v1.Get("/videos", api.FindallVideosHandler)

	g.Get("/health-check", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{"status": "server is up and running", "version": config.ReleaseVersion})
	})
}
