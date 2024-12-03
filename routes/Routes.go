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
}
