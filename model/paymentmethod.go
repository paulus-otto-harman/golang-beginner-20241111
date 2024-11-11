package model

type PaymentMethod struct {
	Name     string `json:"name"`
	Photo    string `json:"photo"`
	IsActive bool   `json:"is_active"`
}
