package controller

import "github.com/gofiber/fiber/v2"

type CreateRolesPayload struct {
	Name string `json:"name"`
}

func CreateRoles(c *fiber.Ctx) error {
	return nil
}
