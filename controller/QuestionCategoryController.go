package controller

import (
	"log"

	"github.com/dickanirwansyah/go-examp/database"
	"github.com/dickanirwansyah/go-examp/model"
	"github.com/dickanirwansyah/go-examp/util"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RequestCategory struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func CreateQuestionCategory(c *fiber.Ctx) error {

	var createCategory RequestCategory

	if err := c.BodyParser(&createCategory); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request payload !")
	}

	var existingQuestionCategory model.QuestionCategory
	if err := database.DB.Where("name = ?", createCategory.Name).First(&existingQuestionCategory).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			questionCategory := model.QuestionCategory{
				Id:        0,
				Name:      createCategory.Name,
				IsDeleted: 0,
			}
			if err := database.DB.Create(&questionCategory).Error; err != nil {
				log.Printf("failed create new question category : %v", err)
				return util.ErrorResponse(c, fiber.StatusInternalServerError, "Something went wrong !")
			} else {
				return util.SuccessResponse(c, fiber.StatusOK, questionCategory)
			}
		} else {
			log.Printf("failed inquiry category by name : %v", err)
			return util.ErrorResponse(c, fiber.StatusInternalServerError, "Something went wrong !")
		}
	}

	return util.ErrorResponse(c, fiber.StatusBadRequest, "Question category name is existing !")
}

func UpdateQuestionCategory(c *fiber.Ctx) error {

	var updateCategory RequestCategory
	if err := c.BodyParser(&updateCategory); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request payload !")
	}

	var existingQuestionCategory model.QuestionCategory
	if err := database.DB.Where("id = ?", updateCategory.Id).First(&existingQuestionCategory).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return util.ErrorResponse(c, fiber.StatusNotFound, "Id question category not found !")
		}
	}

	if existingQuestionCategory.Name != updateCategory.Name {
		if err := database.DB.Where("name = ?", updateCategory.Name).First(&existingQuestionCategory).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				existingQuestionCategory.Name = updateCategory.Name
				if err := database.DB.Model(&existingQuestionCategory).Updates(existingQuestionCategory).Error; err != nil {
					log.Printf("failed updated question category : %v", err)
					return util.ErrorResponse(c, fiber.StatusInternalServerError, "Something went wrong update data question category !")
				} else {
					log.Printf("success update question category : %v", existingQuestionCategory.Id)
					return util.SuccessResponse(c, fiber.StatusOK, existingQuestionCategory)
				}
			}
		} else {
			return util.ErrorResponse(c, fiber.StatusBadRequest, "failed question category name is existing !")
		}
	}
	log.Printf("existing name no update data !")
	return util.SuccessResponse(c, fiber.StatusOK, updateCategory)
}

func GetQuestionCategory(c *fiber.Ctx) error {

	id := c.Params("id")

	categoryQuestion, err := HelperGetQuestionCategoryByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("question category by id not found : %v", id)
			return util.ErrorResponse(c, fiber.StatusNotFound, "question category not found")
		}
		log.Printf("failed get data question category by id :%v", err)
		return util.ErrorResponse(c, fiber.StatusInternalServerError, "something went wrong !")
	}

	return util.SuccessResponse(c, fiber.StatusOK, categoryQuestion)
}

func HelperGetQuestionCategoryByID(id string) (*model.QuestionCategory, error) {
	var existingCategory model.QuestionCategory
	if err := database.DB.First(&existingCategory, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &existingCategory, nil
}
