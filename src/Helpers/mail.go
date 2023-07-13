package Helpers

import (
	"errors"
	"gopkg.in/gomail.v2"
	"log"
	"strconv"
)

func SendMail(toParam string, hashType string, hash string, MailData string) error {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config", err)
	}

	MailAddress := config.MailAddress

	MailPassword := config.MailPassword

	UIURL := config.UIURL

	MailServiceAddress := config.MailServiceAddress

	MailServicePort := config.MailServicePort

	from := MailAddress
	password := MailPassword

	host := MailServiceAddress
	port := MailServicePort
	intPort, err := strconv.Atoi(port)
	if err != nil {
		return err
	}
	subject := ""
	body := ""
	if hashType == "EMAIL VERIFICATION" || hashType == "MEMBER EMAIL VERIFICATION" {
		subject = "Email Verification"
		body = "<p>Hello, kindly use this <a href=\"" + UIURL + "/verify-email/" + hash + "\">link</a> to verify your e-mail.</p>"
	} else if hashType == "PASSWORD RESET" || hashType == "MEMBER PASSWORD RESET" {
		subject = "Password Reset"
		body = "<p>You requested for reset password, kindly use this <a href=\"" + UIURL + "/verify-password/" + hash + "\">link</a> to reset your password.</p>"
	} else if hashType == "EMAIL CHANGE" || hashType == "MEMBER EMAIL CHANGE" {
		subject = "Chane E-mail"
		body = "<p>You requested to change your e-mail to " + MailData + ", kindly use this <a href=\"" + UIURL + "/verify-email/" + hash + "\">link</a> to confirm change.</p>"
	} else {
		return errors.New("Invalid Hash Type")
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", toParam)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	n := gomail.NewDialer(host, intPort, from, password)

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}

	return nil
}
