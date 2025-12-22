package utils

import "github.com/gofiber/fiber/v2/log"

func ErrorResponse(err string) map[string]interface{} {
	log.Error(err)
	return map[string]interface{}{
		"status": "FAILED",
		"error":  err,
	}
}

func SuccessResponse(msg string) map[string]interface{} {
	log.Info(msg)
	resp := map[string]interface{}{
		"status": "SUCCESS",
	}
	if msg != "" {
		resp["message"] = msg
	}
	return resp

}
