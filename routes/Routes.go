package routes

import (
	"github.com/dickanirwansyah/go-examp/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Post("/api/forgot-password", controller.ForgotPassword)
	app.Put("/api/update-password", controller.UpdatePassword)

	/** todo secure endpoint by roles **/
	app.Post("/api/category/create", controller.CreateQuestionCategory)
	app.Put("/api/category/update", controller.UpdateQuestionCategory)
	app.Get("/api/category/find/:id", controller.GetQuestionCategory)
	app.Get("/api/category/list", controller.PageDataQuestionCategory)

	app.Get("/api/account/list", controller.PageAccount)
	app.Post("/api/account/upload", controller.UploadImage)
	app.Static("/api/account/image", "./upload")
	app.Get("/api/account/find/:id", controller.GetAccountById)
}
