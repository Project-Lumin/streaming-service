package api

import (
	"os"
	"path/filepath"
	"streaming-service/models"
	"streaming-service/repo"
	"streaming-service/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateVideoDataHandler(c *fiber.Ctx) error {
	newVideo := models.Video{}
	err := c.BodyParser(&newVideo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse(err.Error()))
	}
	filePath := filepath.Join(StoreDir, newVideo.Id)
	_, err = os.Stat(filePath+".mp4")
	if os.IsNotExist(err) {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse("video for the provided id doesnt exist"))
	}
	err = repo.DB_CreateVideo(&newVideo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse(""))
}
