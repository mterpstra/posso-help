package email

import (
  "fmt"
	"net/smtp"
  "os"
)


const template = `To complete your registration to ZapManejo
enter this code %s on our registration screen.`

func SendRegistrationEmail(to, code string) error {
	smtpHost := "smtp.protonmail.ch"
	smtpPort := "587"
	username := "info@possohelp.com"
	password := os.Getenv("SMTP_PASSWORD")

	from := username
	toEmail := []string{to}

	subject := "Subject: Registation Code to ZapManejo!\r\n"
	body := fmt.Sprintf(template, code)
	msg := []byte(subject + "\r\n" + body)

	auth := smtp.PlainAuth("", username, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, toEmail, msg)
	if err != nil {
    return err
	}

	fmt.Println("Email sent successfully!")
  return nil
}
