package main

import (
	"fmt"
	"log"

	"github.com/dickanirwansyah/go-examp/database"
	"github.com/dickanirwansyah/go-examp/routes"
	"github.com/dickanirwansyah/go-examp/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	//init database
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
	}))

	routes.SetupRoutes(app)

	//provide data generate token & verify token
	testUser := util.DataUser{
		Id:    1,
		Email: "dickanirwansyah@gmail.com",
	}

	//generate token
	token, err := util.GenerateJwt(testUser)
	if err != nil {
		log.Fatalf("Error generate JWT token : %v", err)
	}

	//print token jwt
	fmt.Println("Generate JWT TOKEN : ")
	fmt.Println(token)

	//verify token
	dataUser, claims, err := util.VerifyJwt(token)
	if err != nil {
		log.Fatal("Error verify JWT TOKEN : ", err)
	}

	fmt.Println("Verify JWT TOKEN : ")
	fmt.Println("data_user : ", dataUser.Email)
	fmt.Println("exp : ", claims["exp"])

	//test send email
	err = util.SendEmail("dickanirwansyah@gmail.com", "Application Starting up", "<h1>Application UP !</h1")
	if err != nil {
		log.Fatalf("Could not sent email :%v", err)
	}

	app.Listen(":8000")
}
