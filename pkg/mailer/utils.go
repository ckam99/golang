package mailer

import (
	"os"
	"strconv"
)

func NewMail() *MailMessage {
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	tls, _ := strconv.ParseBool(os.Getenv("MAIL_TLS"))
	return &MailMessage{
		From:           os.Getenv("MAIL_FROM"),
		SMTPHost:       os.Getenv("MAIL_SERVER"),
		SMTPPort:       port,
		SMTPUser:       os.Getenv("MAIL_USER"),
		SMTPPassword:   os.Getenv("MAIL_PASSWORD"),
		TemplateFolder: "./resource/templates",
		TLS:            tls,
	}
}

func Notify(to []string, subject string, data interface{}, template string, cc []ReplyTo) error {
	m := NewMail()
	return m.SendMail(subject, to, template, data, cc)
}
