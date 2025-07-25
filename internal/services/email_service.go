package services

type Emailer interface {
}

type Email struct{}

func NewEmailer() Emailer {
	return &Email{}
}
