package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func sendEmailSimple() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	password := os.Getenv("EMAILPASSWORD")
	auth := smtp.PlainAuth(
		"",
		"charlesemilienpoulin@gmail.com",
		password,
		"smtp.gmail.com",
	)

	msg := "Subject: My special subject\n\nThis is the body of the email."

	erro := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"charlesemilienpoulin@gmail.com",
		[]string{"charlesemilienpoulin@gmail.com"},
		[]byte(msg),
	)

	if erro != nil {
		fmt.Println(err)
	}
}

func main() {
	sendEmailSimple()
}
