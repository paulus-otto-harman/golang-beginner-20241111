package service

import "20241111/repository"

type PaymentMethodService struct {
	PaymentMethodRepo repository.PaymentMethod
}

func InitPaymentMethodService(repo repository.PaymentMethod) *PaymentMethodService {
	return &PaymentMethodService{PaymentMethodRepo: repo}
}
