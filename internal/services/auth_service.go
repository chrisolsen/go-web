package services

type Authenticator interface {
}

type Auth struct{}

func NewAuthenticator() Authenticator {
	return &Auth{}
}
