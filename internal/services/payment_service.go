package services

type PaymentServicer interface {

}

type PaymentService struct{}
func NewPaymentService() PaymentServicer {
	return &PaymentService{}
}