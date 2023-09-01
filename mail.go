package main

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func sendMail(email, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "Cherava <cheravafoss@outlook.com>")
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html",
		fmt.Sprintf("Here are the changes: <br/> <div> <br/> %s <div/>", body))

	d := gomail.NewDialer("smtp-mail.outlook.com", 587, "cheravafoss@outlook.com", "Cherava123")

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
