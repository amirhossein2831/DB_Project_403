package models

import "time"

type Transaction struct {
	ID                   int        `json:"id" sql:"id"`
	Type                 string     `json:"account_id" sql:"account_id"`
	Amount               float64    `json:"amount" sql:"amount"`
	SourceAccountId      int        `json:"description" sql:"source_account_id"`
	DestinationAccountId *int       `json:"type" sql:"destination_account_id"`
	CreatedAt            *time.Time `json:"created_at" sql:"created_at"`
	SourceAccount        *Account   `json:"source_account" sql:"source_account"`
	DestinationAccount   *Account   `json:"destination_account" sql:"-"`
}
