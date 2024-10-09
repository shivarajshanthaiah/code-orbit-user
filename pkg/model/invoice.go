package model

import "gorm.io/gorm"

type Invoice struct {
	gorm.Model

	UserID         string
	UserEmail      string
	SubscriptionID int
	PaymentID      string
	Plan           string
	Amount         uint32
	PaymentStatus  string `gorm:"default:PENDING"`
}

type RazorPay struct {
	InvoiceID       uint32
	RazorPaymentID  string
	RazorPayOrderID string
	Signature       string
	AmountPaid      uint32
}
