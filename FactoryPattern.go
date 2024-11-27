package main

import (
	"fmt"
)

type Notification interface {
	Send(message string)
}

type EmailNotification struct {
	EmailAddress string
}

func (e *EmailNotification) Send(message string) {
	fmt.Printf("Enviando email a %s: %s\n", e.EmailAddress, message)
}

type SMSNotification struct {
	PhoneNumber string
}

func (s *SMSNotification) Send(message string) {
	fmt.Printf("Enviando SMS a %s: %s\n", s.PhoneNumber, message)
}

func NotificationFactory(notificationType string, contactInfo string) Notification {
	switch notificationType {
	case "email":
		return &EmailNotification{EmailAddress: contactInfo}
	case "sms":
		return &SMSNotification{PhoneNumber: contactInfo}
	default:
		return nil
	}
}

func main() {

	email := NotificationFactory("email", "contacto@dominio.com")
	sms := NotificationFactory("sms", "123-456-789")

	if email != nil {
		email.Send("¡Hola, tienes un nuevo mensaje!")
	}

	if sms != nil {
		sms.Send("¡Tu código de verificación es 12345!")
	}

}
