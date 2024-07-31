package service

import "log"

// EmailServiceInterface визначає інтерфейс для служби відправки email
type EmailServiceInterface interface {
	SendOrderEmail(orderID string)
}

// EmailService структура для служби відправки email
type EmailService struct{}

// NewEmailService створює новий екземпляр EmailService
func NewEmailService() *EmailService {
	return &EmailService{}
}

// SendOrderEmail симулює відправку email після замовлення фільму
func (s *EmailService) SendOrderEmail(orderID string) {
	log.Printf("Email sent for order ID: %s", orderID)
}
