package email

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/d11m08y03/CC-EOY/config"
	"github.com/d11m08y03/CC-EOY/logger"
	"github.com/d11m08y03/CC-EOY/models"
	"gopkg.in/mail.v2"
	gomail "gopkg.in/mail.v2"
)

var emails []models.Email
var emailCount int
var index int

var message *mail.Message
var emailContent string

func InitEmails() {
	var err error
	emails, err = models.GetAllEmails()
	if err != nil {
		logger.Fatal("Emails could not be read")
		return
	}

	index = 0
	emailCount = len(emails)

	logger.Info("Emails initialised")

	htmlFilePath := "./email/new-email.html"
	htmlContent, err := os.ReadFile(htmlFilePath)
	if err != nil {
		log.Fatalf("Failed to read HTML file: %v", err)
	}

	emailContent = string(htmlContent)

	message = mail.NewMessage()

	message.Embed("./email/images/facebook2x.png")
	message.Embed("./email/images/instagram2x.png")
	message.Embed("./email/images/linkedin2x.png")
	message.Embed("./email/images/tiktok2x.png")
}

func SendEmail(recipient string, name string) {
	logger.Info("Attempting to send email")
	sender := emails[index]
	index = (index + 1) % (emailCount)

	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	message.SetHeader("From", sender.Email)

	if config.Environment == "Prod" {
		message.SetHeader("To", recipient)
	} else {
		message.SetHeader("To", config.EmailRecipient)
	}

	message.SetHeader("Subject", "EOY Party")

	modifiedEmail := strings.Replace(emailContent, "Students,", fmt.Sprintf("%s,", name), 1)
	message.SetBody("text/html", modifiedEmail)

	d := gomail.NewDialer(smtpHost, smtpPort, sender.Email, sender.AppPassword)
	if err := d.DialAndSend(message); err != nil {
		logger.Error(err.Error())
		return
	}
}
