package models

import "time"

type BankTransactions struct {
	ID                       int       `json:"id" sql:"id"`
	Type                     string    `json:"type" sql:"type"`
	SourceAccountNumber      string    `json:"source_account_number" sql:"source_account_number"`
	DestinationAccountNumber *string   `json:"destination_account_number" sql:"destination_account_number"`
	Amount                   float64   `json:"amount" sql:"amount"`
	CreatedAt                time.Time `json:"created_at" sql:"created_at"`
}
