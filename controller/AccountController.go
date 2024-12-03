package controller

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/dickanirwansyah/go-examp/database"
	"github.com/dickanirwansyah/go-examp/model"
	"github.com/dickanirwansyah/go-examp/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func PageAccount(c *fiber.Ctx) error {

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limitData, _ := strconv.Atoi(c.Query("size", "5"))

	return model.GenericPaginate(c, database.DB, &model.Accounts{}, page, limitData)
}

type PayloadUploadImage struct {
	AccountId          int    `json:"account_id"`
	ImageProfileBase64 string `json:"image_profile_base64"`
}

func UploadImage(c *fiber.Ctx) error {

	var payloadUploadImage PayloadUploadImage
	var currentErrors []string

	if err := c.BodyParser(&payloadUploadImage); err != nil {
		log.Printf("error body parser upload image account %v", err)
		currentErrors = append(currentErrors, "Invalid payload upload !")
	}

	var existingAccount model.Accounts
	if err := database.DB.Where("id = ?", payloadUploadImage.AccountId).First(&existingAccount).Error; err != nil {
		log.Printf("failed check existing account : %v", err)
		if err == gorm.ErrRecordNotFound {
			currentErrors = append(currentErrors, "Invalid account id !")
		} else {
			currentErrors = append(currentErrors, "Something went wrong check account id !")
		}
	}

	if payloadUploadImage.ImageProfileBase64 == "" {
		currentErrors = append(currentErrors, "Invalid image !")
	}

	imageData, err := base64.StdEncoding.DecodeString(payloadUploadImage.ImageProfileBase64)
	if err != nil {
		currentErrors = append(currentErrors, "Failed decode image !")
	}

	if len(currentErrors) > 0 {
		return util.ErrorResponse(c, fiber.StatusBadRequest, currentErrors)
	}

	fileName := uuid.New().String() + ".png"

	uploadDir := "./upload"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			log.Printf("error upload image because %v ", err)
			return util.ErrorResponse(c, fiber.StatusInternalServerError, "Failed upload image")
		}
	}

	filePath := filepath.Join(uploadDir, fileName)

	if err := os.WriteFile(filePath, imageData, 0644); err != nil {
		log.Printf("error write image because %v", err)
		return util.ErrorResponse(c, fiber.StatusInternalServerError, "Failed upload image")
	}

	baseUrl := fmt.Sprintf("http://%s", c.Hostname())
	fileUrl := fmt.Sprintf("%s/api/account/image/%s", baseUrl, fileName)

	existingAccount.ImageProfile = fileUrl
	if err := database.DB.Model(&existingAccount).Updates(existingAccount).Error; err != nil {
		log.Printf("error updates image account by account id %v ", err)
		return util.ErrorResponse(c, fiber.StatusInternalServerError, "Failed upload image !")
	}

	return util.SuccessResponse(c, fiber.StatusOK, fiber.Map{
		"image_url_path":   filePath,
		"image_url_static": fileUrl,
		"account":          existingAccount,
	})
}
