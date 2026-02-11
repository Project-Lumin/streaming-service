package api

import (
	"streaming-service/models"
	"streaming-service/repo"
	"streaming-service/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateUserPrefetchedVideosHandler(c *fiber.Ctx) error {
	prefetchedVideos := models.CreatePrefetchedVideosInput{}
	err := c.BodyParser(&prefetchedVideos)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse(err.Error()))
	}
	date := time.Now().Format("2006-01-02")
	prefetchedVideos.Date = date
	videos, err := repo.DB_CreateUserPrefetchedVideos(&prefetchedVideos)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(videos)
}
