package services

type EmailServicer interface {
}

type EmailService struct{}

func NewEmailService() EmailServicer {
	return &EmailService{}
}