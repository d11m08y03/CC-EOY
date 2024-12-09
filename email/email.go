package email

import (
	"log"
	"os"

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

	// htmlFilePath := "./email/new-email.html" // Path to your HTML file
	// htmlContent, err := os.ReadFile(htmlFilePath)
	// if err != nil {
	// 	log.Fatalf("Failed to read HTML file: %v", err)
	// }

	// message.Embed("./email/images/LOGO-SU.png")
	// message.Embed("./email/images/bigg-frankii-5.jpg")
	// message.Embed("./email/images/Blakkayo.jpg")
	// message.Embed("./email/images/Diferans_Musical_Crew.jpg")
	// message.Embed("./email/images/facebook2x.png")
	// message.Embed("./email/images/Final_final_collage.png")
	// message.Embed("./email/images/instagram2x.png")
	// message.Embed("./email/images/linkedin2x.png")
	// message.Embed("./email/images/MUSIC-HEADER.png")
	// message.Embed("./email/images/output-onlinetools.png")
	// message.Embed("./email/images/the-prophecy-1.jpg")
	// message.Embed("./email/images/tiktok2x.png")
	// message.Embed("./email/images/vinile.png")

	// message.SetBody("text/html", string(htmlContent))
}

func SendEmail(recipient string) {
	logger.Info("Attempting to send email")
	sender := emails[index]
	index = (index + 1) % (emailCount)

	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	htmlFilePath := "./email/new-email.html"
	htmlContent, err := os.ReadFile(htmlFilePath)
	if err != nil {
		logger.Error(err.Error())
	}

	message.SetHeader("From", sender.Email)

	if config.Environment == "Prod" {
		message.SetHeader("To", recipient)
	} else {
		message.SetHeader("To", config.EmailRecipient)
	}

	message.SetHeader("Subject", "EOY Party")
	message.SetBody("text/plain", string(htmlContent))

	d := gomail.NewDialer(smtpHost, smtpPort, sender.Email, sender.AppPassword)
	if err := d.DialAndSend(message); err != nil {
		logger.Error(err.Error())
		return
	}
}

func SendEmailTest() {
	logger.Info("Attempting to send email")

	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	sender := emails[index]
	index = (index + 1) % (emailCount)

	to := "muhammad.kurmally6@umail.uom.ac.mu"

	m := mail.NewMessage()

	sender.Email = "uomcomputerclub1108@gmail.com"
	sender.AppPassword = "vexr fvdz efpm glxx"

	m.SetHeader("From", sender.Email)

	m.SetHeader("To", to)

	m.SetHeader("Subject", "Test Email using Gomail")

	// Set the body of the email (plain text or HTML)
	m.SetBody("text/plain", "Hello, this is a test email sent from Go using Gomail!")

	dialer := mail.NewDialer(smtpHost, smtpPort, sender.Email, sender.AppPassword)

	// Send the email
	if err := dialer.DialAndSend(m); err != nil {
		log.Println(sender.Email)
		log.Println(sender.AppPassword)
		log.Fatalf("Failed to send email: %v", err)
	}

	log.Println("Email sent successfully!")
}

// zakariyyazak43@gmail.com App password
// dwvp fqed zmtf kiwx
