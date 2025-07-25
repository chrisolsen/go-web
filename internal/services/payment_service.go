package services

type Paymenter interface {
}

type Payment struct{}

func NewPaymenter() Paymenter {
	return &Payment{}
}
