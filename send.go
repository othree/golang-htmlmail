package main

import (
	"crypto/tls"
	"io/ioutil"
	"gopkg.in/gomail.v1"
	"github.com/vaughan0/go-ini"
)

func main() {
	config, err := ini.LoadFile("config.ini")
	if err != nil {
		panic("Config file not loaded.")
	}
	USER, ok := config.Get("gmail", "user")
	if !ok {
		panic("Gmail user not set.")
	}
	PASS, ok := config.Get("gmail", "pass")
	if !ok {
		panic("Gmail pass not set.")
	}
	TO, ok := config.Get("to", "account")
	if !ok {
		panic("To account not set.")
	}


	msg := gomail.NewMessage()
	msg.SetHeader("From", USER + "@gmail.com")
	msg.SetHeader("To", TO)

	title, err := ioutil.ReadFile("title")
	if err != nil {
		panic(err)
	}
	html, err := ioutil.ReadFile("body.html")
	if err != nil {
		panic(err)
	}
	msg.SetHeader("Subject", string(title))
	msg.SetBody("text/html", string(html))

	// Send the email to Bob
	mailer := gomail.NewMailer("smtp.gmail.com", USER, PASS, 587, gomail.SetTLSConfig(&tls.Config{InsecureSkipVerify: true}))
	if err := mailer.Send(msg); err != nil {
		panic(err)
	}
}
