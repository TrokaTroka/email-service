package email

import (
	"log"
	"mail-sender/config"
	"mail-sender/models"
	"net/smtp"
)

func getEmailConfig() (addr string, a smtp.Auth, from string) {

	smtpHost, err := config.GetConfig("smtp_host")

	if err != nil {
		log.Fatal(err)
	}
	
	smtpPort, err := config.GetConfig("smtp_port")

	if err != nil {
		log.Fatal(err)
	}

	from, err = config.GetConfig("service_email")

	if err != nil {
		log.Fatal(err)
	}

	password, err := config.GetConfig("service_email_password")

	if err != nil {
		log.Fatal(err)
	}

	a = smtp.PlainAuth("", from, password, smtpHost)

	addr = smtpHost+":"+smtpPort;
	return addr, a, from
}

func SendEmail(email *models.Email) error {

	to := []string{email.To}

	subject := "Subject: "+ email.Subject + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + email.Body)

	address, auth, from := getEmailConfig()
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func SaveEmail(email *models.Email) error {

	err := CreateEmail(email)
	
	if err != nil {
		return err
	}

	return nil
}