package main

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func sendEmailSimpleHTML(subject string, html string, to []string) {

	erro := godotenv.Load()
	password := os.Getenv("EMAILPASSWORD")

	auth := smtp.PlainAuth(
		"",
		"charlesemilienpoulin@gmail.com",
		password,
		"smtp.gmail.com",
	)
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	msg := "Subject:" + subject + "\n" + headers + "\n\n" + html

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
	//sendEmailSimple(
	//"another subject", "another body", []string{"charlesemilienpoulin@gmail.com"}
	//)

	sendEmailSimpleHTML(
		"another subject", "<h1> im a header </h1> <p> im a paragraph </p>", []string{"charlesemilienpoulin@gmail.com"})
}
