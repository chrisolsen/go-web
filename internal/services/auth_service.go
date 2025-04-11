package services

type AuthServicer interface {

}

type AuthService struct{}

func NewAuthService() AuthServicer {
	return &AuthService{}
}