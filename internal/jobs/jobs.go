package jobs

import (
	"log"

	"github.com/ckam225/golang/fiber/pkg/mailer"
)

func RegisterNotificationChannel() {
	go func() {
		for notify := range mailer.NotificationChannel {
			err := mailer.Notify(notify.To, notify.Subject, notify.Data, notify.Template, notify.CC)
			if err != nil {
				log.Println(err.Error())
			} else {
				log.Println("Email sent successfully", notify.To)
			}
		}
	}()
}

func UnregisterNotificationChannel() {
	defer close(mailer.NotificationChannel)
}
