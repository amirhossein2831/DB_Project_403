package models

import "time"

type AccountType string
type AccountStatus string

const (
	SavingsAccountType  AccountType = "savings"
	CurrentAccountType  AccountType = "current"
	BusinessAccountType AccountType = "business"
)

const (
	ActiveAccountStatus  AccountStatus = "active"
	ClosedAccountStatus  AccountStatus = "closed"
	PendingAccountStatus AccountStatus = "pending"
)

type Account struct {
	ID            int           `json:"id" sql:"id"`
	AccountNumber string        `json:"account_number" sql:"account_number"`
	Type          AccountType   `json:"type" sql:"type"`
	Amount        float64       `json:"amount" sql:"amount"`
	Status        AccountStatus `json:"status" sql:"status"`
	CustomerID    int           `json:"customer_id" sql:"customer_id"`
	Customer      *Customer     `json:"customer" sql:""`
	CreatedAt     time.Time     `json:"created_at" sql:"created_at"`
	ClosedAt      time.Time     `json:"closed_at" sql:"closed_at"`
}
