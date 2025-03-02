package controllers

import (
	"log"
	"net/smtp"
)

type EmailController struct {
	sender    string
	password  string
	smtp_host string
	smtp_port string
}

func NewEmailController(sender, password, smtp_host, smtp_port string) EmailController {
	return EmailController{
		sender:    sender,
		password:  password,
		smtp_host: "smtp.gmail.com",
		smtp_port: smtp_port,
	}
}

func (ec EmailController) SendEmail(to, subject, body string) {

	mensaje := []byte(subject + "\r\n" + body)

	// Configurar la autenticación SMTP
	auth := smtp.PlainAuth("", ec.sender, ec.password, ec.smtp_host)

	// Enviar el correo
	err := smtp.SendMail(ec.smtp_host+":"+ec.smtp_port, auth, ec.sender, []string{to}, mensaje)
	if err != nil {
		log.Printf("Error al enviar el email: %v", err)
		return
	}
}
