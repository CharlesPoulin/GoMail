package main

import (
	"bytes"
	"fmt"
	"gopkg.in/gomail.v2"
	"html/template"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func sendEmailSimpleHTML(subject string, templatePath string, to []string) {

	var body bytes.Buffer

	t, err := template.ParseFiles(templatePath)

	t.Execute(&body, struct{ Name string }{Name: "Charles"})

	err = godotenv.Load()
	password := os.Getenv("EMAILPASSWORD")

	auth := smtp.PlainAuth(
		"",
		"charlesemilienpoulin@gmail.com",
		password,
		"smtp.gmail.com",
	)
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	msg := "Subject:" + subject + "\n" + headers + "\n\n" + body.String()

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"charlesemilienpoulin@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func sendEmailSimple(subject string, body string, to []string) {

	err := godotenv.Load()
	password := os.Getenv("EMAILPASSWORD")

	auth := smtp.PlainAuth(
		"",
		"charlesemilienpoulin@gmail.com",
		password,
		"smtp.gmail.com",
	)

	msg := "Subject:" + subject + "\n" + body

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"charlesemilienpoulin@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func sendGoMail(templatePath string) {

	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	t.Execute(&body, struct{ Name string }{Name: "Charles"})
	err = godotenv.Load()
	password := os.Getenv("EMAILPASSWORD")

	if err != nil {
		fmt.Println(err)
		return
	}
	// send with gomail

	m := gomail.NewMessage()
	m.SetHeader("From", "charlesemilienpoulin@gmail.com")
	m.SetHeader("To", "charlesemilienpoulin@gmail.com", "charlespoulin@hotmail.ca")
	m.SetAddressHeader("Cc", "charlesemilienpoulin@gmail.com", "charles")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", body.String())
	m.Attach("./4x.webp")

	d := gomail.NewDialer("smtp.gmail.com", 587, "charlesemilienpoulin@gmail.com", password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
func main() {
	//sendEmailSimple(
	//"another subject", "another body", []string{"charlesemilienpoulin@gmail.com"}
	//)

	//sendEmailSimpleHTML(
	//	"another subject", "./test.html", []string{"charlesemilienpoulin@gmail.com"})

	sendGoMail("./test.html")
}
