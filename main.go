package main

import (
	"flag"
	"github.com/adelolmo/notify-mail/mail"
	"log"
	"os"
	"strings"
)

func main() {
	notifyMail, err := mail.NewNotification()
	if err != nil {
		log.Fatal(err)
	}
	recipient, subject, message, template, variables := parseArguments()
	if len(message) > 0 {
		err = notifyMail.Send(recipient, subject, message)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = notifyMail.SendTemplate(recipient, subject, template, variables)
	if err != nil {
		log.Fatal(err)
	}
}

func parseArguments() (string, string, string, string, map[string]string) {
	recipient := flag.String("recipient", "", "Recipient address")
	subject := flag.String("subject", "", "Notification subject")
	message := flag.String("message", "", "Message content")
	template := flag.String("template", "", "Template filename")
	variablesString := flag.String("variables", "",
		"Template variables with format {{var1}}=value of var1,{{var2}}=value of var2")
	flag.Parse()

	if len(*recipient) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	variables := strings.Split(*variablesString, ",")
	m := make(map[string]string)
	if len(*variablesString) > 0 {
		for _, variable := range variables {
			keyValue := strings.Split(variable, "=")
			m[keyValue[0]] = keyValue[1]
		}
	}
	return *recipient, *subject, *message, *template, m
}
