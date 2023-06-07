package email

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

const (
	senderName = "Automated DHC"
	smtpAuthAddress = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type Contact struct {
	Name  string 	`json:"name"`
	Email string 	`json:"email"`
	Phone string 	`json:"phone"`
}

func Send(contact Contact) error {
	gmailAppKey := os.Getenv("GMAIL_APP_KEY")
	gmailSender := os.Getenv("GMAIL_USER")

	msgString := fmt.Sprintf("Nome: %s\nEmail: %s\nPhone: %s", contact.Name, contact.Email, contact.Phone)

	emailMessage := &email.Email {
		To: []string{gmailSender},
		From: "Automated <" + gmailSender + ">",
		Subject: "Cadastro",
		Text: []byte(msgString),
	}

	auth := smtp.PlainAuth(senderName, gmailSender, gmailAppKey, smtpAuthAddress)
	return emailMessage.Send(smtpServerAddress, auth)
}