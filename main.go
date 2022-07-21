package main

import (
	"fmt"

	"example.com/go-mail/helpers"
	email "example.com/go-mail/mail"
	"gopkg.in/mail.v2"
)

func main() {
	SendCustomMail()
	// SendMail()
}

func SendMail() {
	m := mail.NewMessage()
	m.SetHeader("From", "alex@example.com")
	m.SetHeader("To", "bob@example.com", "cora@example.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	buff, err := helpers.ParseTemplate("./templates/test_mail.html", map[string]string{
		"name":  "Martin Ballo",
		"email": "info@example.com",
	})
	if err != nil {
		panic(err)
	}

	m.SetBody("text/html", buff)
	m.Attach("./iphone-14.jpg")

	// d := mail.NewDialer("smtp.example.com", 587, "user", "123456")
	d := mail.NewDialer("smtp.mailtrap.io", 2525, "f10c9fcb8036ed", "db94187020064b")

	d.StartTLSPolicy = mail.MandatoryStartTLS

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err.Error())
		panic(err)
	} else {
		fmt.Println("Email successfully sent!")
	}
}

func SendCustomMail() {
	username := "f10c9fcb8036ed"
	password := "db94187020064b"
	host := "smtp.mailtrap.io"
	port := "2525"

	templateData := struct {
		Name string
		URL  string
	}{
		Name: "Dhanush",
		URL:  "http://geektrust.in",
	}
	to := []string{"junk@junk.com"}

	m := email.MailMessage{
		From:    "test@example.com",
		To:      to,
		Subject: "Hello Junk!",
		Body:    "Hello, World!",
	}
	config := &email.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
	ok, err := m.SendHTML(config, "./templates/mail.html", templateData)
	if err != nil {
		panic(err)
	}
	fmt.Println(ok)
}
