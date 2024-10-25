package email

import (
	"fmt"
	"newsletter-service/pkg/templates"

	"gopkg.in/gomail.v2"
)

func SendEmail(server *gomail.Dialer, title, attachment, html, recipient, unsuscribeAllUrl, unsuscribyCategoryUrl string) error {
	m := gomail.NewMessage()

	if html == "" {
		html = templates.DefaultTemplate
	}

	m.SetHeader("From", "example@gmail.com")
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", title)

	footer := fmt.Sprintf(templates.FooterTemplate, unsuscribeAllUrl, unsuscribyCategoryUrl)

	m.SetBody("text/html", html+footer)
	m.Attach(attachment)

	if err := server.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
