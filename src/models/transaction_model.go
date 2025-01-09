package models

import "time"

type Transaction struct {
	ID                   int        `json:"id" sql:"id"`
	Type                 string     `json:"type" sql:"account_id"`
	Amount               float64    `json:"amount" sql:"amount"`
	SourceAccountId      int        `json:"source_account_id" sql:"source_account_id"`
	DestinationAccountId *int       `json:"destination_account_id" sql:"destination_account_id"`
	CreatedAt            *time.Time `json:"created_at" sql:"created_at"`
	SourceAccount        *Account   `json:"source_account" sql:""`
	DestinationAccount   *Account   `json:"destination_account" sql:"-"`
}
