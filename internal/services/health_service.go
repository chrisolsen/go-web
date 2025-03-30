package services

type HealthServicer interface {
	Check() string
}

type HealthService struct{}

func NewHealthService() HealthServicer {
	return &HealthService{}
}

func (hs *HealthService) Check() string {
	return "OK"
}
