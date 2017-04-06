package mail

import (
	"net/smtp"
	"fmt"
	"log"
	"os"
)

type Mail struct {
	Sender         string
	Authentication smtp.Auth
}

func NewNotification() (*Mail, error) {
	keys := []string{
		"NOTIFY_MAIL_SENDER",
		"NOTIFY_MAIL_PASSWORD",
	}
	n := map[string]string{}
	for _, key := range keys {
		v := os.Getenv(key)
		if v == "" {
			return nil, fmt.Errorf("environment variable %q is required", key)
		}
		n[key] = v
	}
	auth := smtp.PlainAuth(
		"Notification",
		n["NOTIFY_MAIL_SENDER"],
		n["NOTIFY_MAIL_PASSWORD"],
		"smtp.gmail.com",
	)
	return &Mail{Authentication:auth}, nil
}

func (m *Mail) Send(recipient, subject, message string) {
	header := make(map[string]string)
	header["From"] = m.Sender
	header["To"] = recipient
	header["Subject"] = subject

	content := ""
	for k, v := range header {
		content += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	content += "\r\n" + message

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		m.Authentication,
		m.Sender,
		[]string{recipient},
		[]byte(content),
	)
	if err != nil {
		log.Fatal(err)
	}
}