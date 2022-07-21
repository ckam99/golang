package mail

import (
	"fmt"
	"net/smtp"

	"example.com/go-mail/helpers"
)

//Request struct
type MailMessage struct {
	From       string
	To         []string
	Subject    string
	Body       string
	ContenType string
}

type SMTPConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

func (r *MailMessage) Send(config *SMTPConfig) (bool, error) {
	if r.ContenType == "" {
		r.ContenType = "text/plain"
	}
	mime := "MIME-version: 1.0;\nContent-Type:" + r.ContenType + "; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.Subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.Body)
	addr := config.Host + ":" + config.Port
	auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)
	if err := smtp.SendMail(addr, auth, r.From, r.To, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *MailMessage) SendHTML(config *SMTPConfig, templateName string, data interface{}) (bool, error) {
	buffer, err := helpers.ParseTemplate(templateName, data)
	if err != nil {
		return false, err
	}
	println("from", r.From)
	fmt.Println("to", r.To)
	r.Body = buffer
	r.ContenType = "text/html"
	return r.Send(config)
}
