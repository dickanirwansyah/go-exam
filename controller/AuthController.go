package controller

import (
	"log"
	"os"
	"time"

	"github.com/dickanirwansyah/go-examp/database"
	"github.com/dickanirwansyah/go-examp/model"
	"github.com/dickanirwansyah/go-examp/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type PayloadRegister struct {
	Id              int    `json:"id"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	FullName        string `json:"full_name"`
	PhoneNumber     string `json:"phone_number"`
	RolesId         uint   `json:"roles_id"`
	AddressDetail   string `json:"address_detail"`
}

type PayloadLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PayloadForgotPassword struct {
	Email string `json:"email"`
}

type PayloadUpdatePassword struct {
	Token           string `json:"token"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func UpdatePassword(c *fiber.Ctx) error {

	var payloadUpdatePassword PayloadUpdatePassword
	var currentErrors []string

	if err := c.BodyParser(&payloadUpdatePassword); err != nil {
		currentErrors = append(currentErrors, "Invalid payload request !")
	}

	if payloadUpdatePassword.Password != payloadUpdatePassword.ConfirmPassword {
		currentErrors = append(currentErrors, "Confirm password do not match with password !")
	}

	var findResetToken model.ResetToken
	if err := database.DB.Where("token=? and is_executed=?", payloadUpdatePassword.Token, "N").First(&findResetToken).Error; err != nil {
		log.Printf("Invalid token %v not exists !", payloadUpdatePassword.Token)
		currentErrors = append(currentErrors, "Invalid token !")
	}

	if len(currentErrors) > 0 {
		return util.ErrorResponse(c, fiber.StatusBadRequest, currentErrors)
	}

	var findCurrentAccount model.Accounts
	if err := database.DB.Where("id = ?", findResetToken.AccountId).First(&findCurrentAccount).Error; err != nil {
		log.Printf("Error account id %v not found !", findResetToken.AccountId)
		return util.ErrorResponse(c, fiber.StatusInternalServerError, "Please contact administrator")
	}

	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payloadUpdatePassword.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Printf("Error hashing password :%v", err)
		return util.ErrorResponse(c, fiber.StatusInternalServerError, "Something went wrong !")
	}

	account := model.Accounts{
		Id:       findCurrentAccount.Id,
		Password: string(hashedPassword),
	}

	//update account & update reset token
	if err := database.DB.Model(&account).Updates(account).Error; err != nil {
		log.Printf("Error update password :%v", err)
		return util.ErrorResponse(c, fiber.StatusInternalServerError, "Failed updated password !")
	}

	resetToken := model.ResetToken{
		Id:         findResetToken.Id,
		IsExecuted: "Y",
	}

	if err := database.DB.Model(&resetToken).Updates(resetToken).Error; err != nil {
		log.Printf("Error update reset token :%v", err)
		return util.ErrorResponse(c, fiber.StatusInsufficientStorage, "Failed updated password")
	}

	return util.SuccessResponse(c, fiber.StatusOK, fiber.Map{
		"account_id":             findCurrentAccount.Id,
		"email":                  findCurrentAccount.Email,
		"status_update_password": "successfully !",
	})
}

func ForgotPassword(c *fiber.Ctx) error {

	var payloadForgotPassword PayloadForgotPassword
	var currentErrors []string

	err := godotenv.Load()
	if err != nil {
		log.Printf("failed load .env for forgot password !")
	}
	urlForgotPassword := os.Getenv("GO_MAIL_FORGOT_PASSWORD")

	if err := c.BodyParser(&payloadForgotPassword); err != nil {
		log.Printf("Error parsing body forgot password : %v", err)
		currentErrors = append(currentErrors, "Invalid request forgot password !")
	}

	if payloadForgotPassword.Email == "" {
		log.Printf("Error, Email cannot be null !")
		currentErrors = append(currentErrors, "Error, Email cannot be empty !")
	}

	var account model.Accounts
	if err := database.DB.Where("email = ? AND is_deleted = ?", payloadForgotPassword.Email, 0).First(&account).Error; err != nil {
		log.Printf("Email not found %v", payloadForgotPassword.Email)
		currentErrors = append(currentErrors, "Error, Email invalid !")
	}

	if len(currentErrors) > 0 {
		log.Printf("Failed forgot password because current errors is not empty !")
		return util.ErrorResponse(c, fiber.StatusBadRequest, currentErrors)
	}

	var findResetToken model.ResetToken
	if err := database.DB.Where("email = ? AND is_executed = ?", payloadForgotPassword.Email, "N").First(&findResetToken).Error; err != nil {
		//create token
		resetToken := model.ResetToken{
			AccountId:  account.Id,
			Email:      account.Email,
			IsExecuted: "N",
			Expires:    time.Now().Add(1 * time.Hour), //1 hours
			Token:      uuid.New().String(),
		}

		if err := database.DB.Create(&resetToken).Error; err != nil {
			log.Printf("failed save reset token : %v", err)
			return util.ErrorResponse(c, fiber.StatusInternalServerError, "Failed reset token !")
		}

		//if success do send email
		fullForgotPasswordUrl := urlForgotPassword + resetToken.Token

		util.SendEmail(account.Email, "Forgot Password", fullForgotPasswordUrl)

		return util.SuccessResponse(c, fiber.StatusOK, fiber.Map{
			"url_forgot_password": fullForgotPasswordUrl,
		})
	}

	if findResetToken.IsExecuted == "N" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"status":  "forgo password failed, please check your email again !",
		})
	}

	return nil
}

func Login(c *fiber.Ctx) error {

	var payloadLogin PayloadLogin
	var currentErrors []string

	if err := c.BodyParser(&payloadLogin); err != nil {
		log.Printf("Error parsing body login : %v", err)
		currentErrors = append(currentErrors, "Invalid payload login")
	}

	if payloadLogin.Email == "" && payloadLogin.Password == "" {
		log.Printf("Error email or password cannot be null")
		currentErrors = append(currentErrors, "Email or Password cannot be null !")
	}

	var account model.Accounts
	if err := database.DB.Where("email = ? and is_deleted = ?", payloadLogin.Email, 0).First(&account).Error; err != nil {
		currentErrors = append(currentErrors, "Invalid payload login !")
	}

	if !util.CheckHashBcryptPassword(payloadLogin.Password, account.Password) {
		log.Printf("Invalid password for email : %v", payloadLogin.Email)
		currentErrors = append(currentErrors, "Invalid email or password")
	}

	generateToken, err := util.GenerateJwt(util.DataUser{
		Id:    int(account.Id),
		Email: account.Email,
	})

	if err != nil {
		log.Printf("Failed to generate JWT token : %v", err)
		currentErrors = append(currentErrors, "Failed generate token")
	}

	if len(currentErrors) > 0 {
		return util.ErrorResponse(c, fiber.StatusBadRequest, currentErrors)
	}

	return util.SuccessResponse(c, fiber.StatusOK, fiber.Map{
		"token":     generateToken,
		"full_name": account.FullName,
	})
}

func Register(c *fiber.Ctx) error {

	var payloadRegister PayloadRegister
	var currentErrors []string

	if err := c.BodyParser(&payloadRegister); err != nil {
		log.Printf("Error parsing body : %v", err)
		currentErrors = append(currentErrors, "Invalid body request !")
	}

	//validation password & confirm password
	if payloadRegister.ConfirmPassword != payloadRegister.Password {
		currentErrors = append(currentErrors, "Password do not match !")
	}

	//check roles
	var roles model.Roles
	if err := database.DB.First(&roles, "id = ? AND is_deleted = 0", payloadRegister.RolesId).Error; err != nil {
		log.Printf("Error finding roles : %v", err)
		currentErrors = append(currentErrors, "Roles id not found !")
	}

	//check current error not emty
	if len(currentErrors) > 0 {
		return util.ErrorResponse(c, fiber.StatusBadRequest, currentErrors)
	}

	//hash password before store it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payloadRegister.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password : %v", err)
		return util.ErrorResponse(c, fiber.StatusInternalServerError, "Something went wrong !")
	}

	user := model.Accounts{
		Email:         payloadRegister.Email,
		FullName:      payloadRegister.FullName,
		PhoneNumber:   payloadRegister.PhoneNumber,
		RolesId:       roles.Id,
		RolesName:     roles.Name,
		AddressDetail: payloadRegister.AddressDetail,
		Password:      string(hashedPassword),
		CreatedAt:     time.Now(),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		log.Printf("Error creating user : %v", err)
		return util.ErrorResponse(c, fiber.StatusInternalServerError, "Something went wrong and failed create new user !")
	}

	return util.SuccessResponse(c, fiber.StatusOK, user)
}
