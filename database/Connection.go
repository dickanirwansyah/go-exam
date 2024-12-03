package database

import (
	"fmt"
	"log"
	"os"

	"github.com/dickanirwansyah/go-examp/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Error failed file .env : %v, please check folder !", err)
	}

	//get environment from file .env
	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	name := os.Getenv("DATABASE_NAME")
	port := os.Getenv("DATABASE_PORT")
	sslmode := os.Getenv("DATABASE_SSL_MODE")

	databaseConnection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, name, port, sslmode)
	database, err := gorm.Open(postgres.Open(databaseConnection), &gorm.Config{})

	if err != nil {
		panic("Error fail connected database !")
	}

	DB = database
	err = database.AutoMigrate(
		&model.Roles{},
		&model.Accounts{},
		&model.ResetToken{},
		&model.Permissions{},
		&model.PermissionsRoles{},
		&model.QuestionCategory{},
		&model.Questions{},
		&model.Answer{})
	if err != nil {
		log.Fatalf("Error running migration : %v", err)
	}
}
