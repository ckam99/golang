package mailer

import (
	"example/fiber/utils"

	mail "gopkg.in/mail.v2"
)

type MailMessage struct {
	From           string
	ContenType     string
	TemplateFolder string
	SMTPHost       string
	SMTPPort       int
	SMTPUser       string
	SMTPPassword   string
}

type ReplyTo struct {
	Email string
	Name  string
}

func MailBuilder(msg *MailMessage, subject string, body string, to []string, cc []ReplyTo) *mail.Message {
	m := mail.NewMessage()
	m.SetHeader("From", msg.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	if msg.ContenType == "" {
		msg.ContenType = "text/plain"
	}
	m.SetBody(msg.ContenType, body)
	if len(cc) > 0 {
		for _, c := range cc {
			m.SetAddressHeader("Cc", c.Email, c.Name)
		}
	}
	return m
}

func (sm *MailMessage) SendMail(
	subject string,
	to []string,
	templateName string,
	body interface{},
	cc []ReplyTo) error {
	m := MailBuilder(sm, subject, "", to, cc)
	// m.Attach("./iphone-14.jpg")
	buff, err := utils.ParseTemplate(sm.TemplateFolder+"/"+templateName, body)
	if err != nil {
		return err
	}
	m.SetBody("text/html", buff)
	d := mail.NewDialer(sm.SMTPHost, sm.SMTPPort, sm.SMTPUser, sm.SMTPPassword)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	err = d.DialAndSend(m)
	return err
}

func (sm *MailMessage) Send(
	subject string,
	body string,
	to []string,
	templateName string,
	templateData interface{},
	cc []ReplyTo) error {
	m := MailBuilder(sm, subject, body, to, cc)
	d := mail.NewDialer(sm.SMTPHost, sm.SMTPPort, sm.SMTPUser, sm.SMTPPassword)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	err := d.DialAndSend(m)
	return err
}
