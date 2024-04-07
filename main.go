package main

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func sendEmailSimple(subject string, body string, to []string) {

	erro := godotenv.Load()
	password := os.Getenv("EMAILPASSWORD")

	auth := smtp.PlainAuth(
		"",
		"charlesemilienpoulin@gmail.com",
		password,
		"smtp.gmail.com",
	)

	msg := "Subject:" + subject + "\n" + body

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"charlesemilienpoulin@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
	if erro != nil {
		fmt.Println(erro)
	}
}

func main() {
	sendEmailSimple("another subject", "another body", []string{"charlesemilienpoulin@gmail.com"})
}
