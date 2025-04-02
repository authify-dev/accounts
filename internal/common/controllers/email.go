package controllers

import (
	"fmt"
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
	// Construir los encabezados del email
	headers := make(map[string]string)
	headers["From"] = ec.sender
	headers["To"] = to
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = `text/html; charset="UTF-8"`

	// Crear el mensaje concatenando los headers y el body
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Convertir el mensaje a []byte
	msg := []byte(message)

	// Configurar la autenticación SMTP
	auth := smtp.PlainAuth("", ec.sender, ec.password, ec.smtp_host)

	// Enviar el correo
	err := smtp.SendMail(ec.smtp_host+":"+ec.smtp_port, auth, ec.sender, []string{to}, msg)
	if err != nil {
		log.Printf("Error al enviar el email: %v", err)
		return
	}
}
