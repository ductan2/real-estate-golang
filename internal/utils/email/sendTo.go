package mail

import (
	"crypto/tls"
	"ecommerce/global"
	"fmt"
	"strings"

	"gopkg.in/gomail.v2"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress `json:"from"`
	To      []string     `json:"to"`
	Subject string       `json:"subject"`
	Body    string       `json:"body"`
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

// Email body
func SendEmailOtp(to []string, from string, otp int) error {
	subject := "Your OTP Code"
	body := fmt.Sprintf("Your OTP code is: %d", otp)

	// Create a new message
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// Create a new SMTP dialer
	d := gomail.NewDialer(global.Config.SMTP.Host, global.Config.SMTP.Port, global.Config.SMTP.Username, global.Config.SMTP.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func SendEmail(to []string, from string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)
	
	d := gomail.NewDialer(global.Config.SMTP.Host, global.Config.SMTP.Port, global.Config.SMTP.Username, global.Config.SMTP.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}


