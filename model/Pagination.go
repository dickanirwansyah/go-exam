package model

import (
	"math"

	"github.com/dickanirwansyah/go-examp/util"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GenericPaginate(c *fiber.Ctx, db *gorm.DB, entityPagination EntityPagination, page int, limitData int) error {

	limit := limitData
	offset := (page - 1) * limit

	pageData := entityPagination.Grab(db, limit, offset)
	totalData := entityPagination.Count(db)

	return util.SuccessResponse(c, fiber.StatusOK, fiber.Map{
		"meta": fiber.Map{
			"total":     totalData,
			"content":   pageData,
			"last_page": math.Ceil(float64(int(totalData) / limit)),
		},
	})
}
