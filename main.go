package main

import (
	"flag"
	"os"
	"log"
	"github.com/adelolmo/notify-mail/mail"
)

func main() {
	notifyMail, err := mail.NewNotification()
	if err != nil {
		log.Fatal(err)
	}
	notifyMail.Send(parseArguments())
}

func parseArguments() (string, string, string) {
	recipient := flag.String("recipient", "", "Recipient address")
	subject := flag.String("subject", "", "Notification subject")
	message := flag.String("message", "", "Message content")
	flag.Parse()

	if len(*recipient) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	return *recipient, *subject, *message
}
