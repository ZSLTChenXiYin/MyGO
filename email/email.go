package email

import (
	"github.com/ZSLTChenXiYin/MyGO/configure"
	"gopkg.in/gomail.v2"
)

type EmailSender struct {
	dialer *gomail.Dialer

	from    string
	subject string
}

func NewEmailSender(conf configure.Configuration, from string, subject string) *EmailSender {
	return &EmailSender{
		dialer: gomail.NewDialer(
			conf.Email().Host(),
			conf.Email().Port(),
			conf.Email().Email(),
			conf.Email().Password(),
		),
		from:    from,
		subject: subject,
	}
}

func (es *EmailSender) Send(to string, content string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", es.from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", es.subject)
	msg.SetBody("text/html", content)

	return es.dialer.DialAndSend(msg)
}
