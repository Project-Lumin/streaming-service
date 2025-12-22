package api

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"streaming-service/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UploadVideoHandler(c *fiber.Ctx) error {
	var (
		fileId = uuid.New().String()
	)
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse(err.Error()))
	}
	files := form.File["video"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse("no files found"))
	}
	if len(files) > 1 {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse("upload only 1 file"))
	}
	file := files[0]

	if err := os.MkdirAll(StoreDir, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse("internal server error"))
	}

	filePath := filepath.Join(StoreDir, fileId+".mp4")
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse(err.Error()))
	}
	dashFolder := "./videos/" + fileId
	if err := os.MkdirAll(dashFolder, 0755); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse(err.Error()))
	}

	dashCmd := exec.Command("ffmpeg", "-i", filePath,
		"-c:v", "libx264", "-c:a", "aac",
		"-f", "dash",
		"-seg_duration", "5",
		"-use_template", "1",
		"-use_timeline", "1",
		dashFolder+"/manifest.mpd",
	)
	if err := dashCmd.Run(); err != nil {
		log.Println("DASH transcoding failed:", err, "[file_id] - ", fileId)
	} else {
		log.Println("DASH transcoding completed", "[file_id] - ", fileId)
	}

	err = os.Remove(filePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse(err.Error()))
	}
	return c.Status(fiber.StatusCreated).JSON(utils.SuccessResponse(fileId))
}
