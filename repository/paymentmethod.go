package repository

import "database/sql"

type PaymentMethod struct {
	Db *sql.DB
}

func InitPaymentMethodRepo(db *sql.DB) *PaymentMethod {
	return &PaymentMethod{db}
}
