package models

type CustomerAccounts struct {
	FirstName     string      `json:"first_name" sql:"first_name"`
	LastName      string      `json:"last_name" sql:"last_name"`
	Phone         string      `json:"phone" sql:"phone"`
	AccountNumber string      `json:"account_number" sql:"account_number"`
	AccountType   AccountType `json:"account_type" sql:"account_type"`
	AccountAmount float64     `json:"account_amount" sql:"account_amount"`
}
