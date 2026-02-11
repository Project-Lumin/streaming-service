package api

import (
	"os"
	"path/filepath"
	"streaming-service/models"
	"streaming-service/repo"
	"streaming-service/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateVideoDataHandler(c *fiber.Ctx) error {
	newVideo := models.Video{}
	err := c.BodyParser(&newVideo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse(err.Error()))
	}
	filePath := filepath.Join(StoreDir, newVideo.FileId)
	_, err = os.Stat(filePath + ".mp4")
	if os.IsNotExist(err) {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse("video for the provided id doesnt exist"))
	}
	date := time.Now().Format("2006-01-02")
	newVideo.DatePosted = date
	err = repo.DB_CreateVideo(&newVideo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse(""))
}

// CreateBulkVideoDataHandler ingests a JSON array of videos and upserts them individually.
func CreateBulkVideoDataHandler(c *fiber.Ctx) error {
	var videos []models.Video
	if err := c.BodyParser(&videos); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse(err.Error()))
	}

	if len(videos) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse("no videos provided"))
	}

	date := time.Now().Format("2006-01-02")
	for i := range videos {
		video := &videos[i]
		if video.FileId == "" || video.Id == "" {
			return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse("video must include id and file_id"))
		}

		filePath := filepath.Join(StoreDir, video.FileId)
		if _, err := os.Stat(filePath + ".mp4"); os.IsNotExist(err) {
			return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse("video for the provided id doesnt exist"))
		} else if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse(err.Error()))
		}

		video.DatePosted = date
	}

	for i := range videos {
		if err := repo.DB_CreateVideo(&videos[i]); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse(err.Error()))
		}
	}

	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse(""))
}
