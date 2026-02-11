package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func FetchSegmentHandler(c *fiber.Ctx) error {
	 videoId := c.Params("videoId")
    file := c.Params("*")

    path := fmt.Sprintf("./videos/%s/%s", videoId, file)
    return c.SendFile(path)
}
