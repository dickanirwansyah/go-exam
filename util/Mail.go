package util

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) error {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Error failed load .env file for send email : %v", err)
	}

	senderEmail := os.Getenv("GO_SENDER_EMAIL")
	senderEmailCredential := os.Getenv("GO_USER_MAIL")
	senderEmailPassword := os.Getenv("GO_PASSWORD_MAIL")
	smtpHost := os.Getenv("GO_SMPT_HOST_MAIL")
	smtpPort := os.Getenv("GO_SMPT_PORT_MAIL")
	numPort, err := strconv.Atoi(smtpPort)

	if err != nil {
		log.Fatalf("failed convert port to string : %v", err)
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", senderEmail)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(smtpHost, numPort, senderEmailCredential, senderEmailPassword)

	if err := dialer.DialAndSend(mailer); err != nil {
		log.Printf("Failed to send email : %v", err)
		return err
	}

	log.Printf("Email sent sucessfully !")
	return nil
}
