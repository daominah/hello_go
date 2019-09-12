package main

import (
	"fmt"

	"gopkg.in/gomail.v1"
)

func main() {
	// need to set gmail "Less secure app access" to "ON"
	mailer := gomail.NewMailer(
		"smtp.gmail.com", "daominah@gmail.com", "talaTung208", 587)
	msg := gomail.NewMessage()
	msg.SetHeader("From", "daominah@gmail.com")
	msg.SetHeader("To", "daominah@gmail.com")
	msg.SetHeader("Subject", "Test sending email")
	msg.SetBody("text/plain", "content")
	err := mailer.Send(msg)
	if err != nil {
		fmt.Println("FAIL", err)
	} else {
		fmt.Println("OK")
	}
}
