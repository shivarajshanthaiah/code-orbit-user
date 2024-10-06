package model

import "gorm.io/gorm"

type Invoice struct {
	gorm.Model

	UserID         string
	UserEmail      string
	SubscriptionID int
	PaymentID      string
	Plan           string
	Amount         float64
	PaymentStatus  string `gorm:"default:PENDING"`
}

type RazorPay struct {
	UserID          string
	RazorPaymentID  string
	RazorPayOrderID string
	Signature       string
	AmountPaid      float64
}
