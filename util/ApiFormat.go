package util

import "github.com/gofiber/fiber/v2"

type ApiResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(ApiResponse{
		Message: "success",
		Status:  status,
		Data:    data,
	})
}

func ErrorResponse(c *fiber.Ctx, status int, errors interface{}) error {
	return c.Status(status).JSON(ApiResponse{
		Message: "error",
		Status:  status,
		Errors:  errors,
	})
}
