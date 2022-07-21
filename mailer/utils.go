package mailer

import (
	"os"
	"strconv"
)

func GetMailInstance() *MailMessage {
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	return &MailMessage{
		From:           os.Getenv("MAIL_FROM"),
		SMTPHost:       os.Getenv("MAIL_SERVER"),
		SMTPPort:       port,
		SMTPUser:       os.Getenv("MAIL_USER"),
		SMTPPassword:   os.Getenv("MAIL_PASSWORD"),
		TemplateFolder: "./resource/templates",
	}
}
