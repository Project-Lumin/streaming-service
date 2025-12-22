package api

import (
	"streaming-service/repo"
	"streaming-service/utils"

	"github.com/gofiber/fiber/v2"
)

func FindallVideosHandler(c *fiber.Ctx) error {
	id := c.Query("id")
	videos, err := repo.DB_FindallVideo(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(videos)
}
