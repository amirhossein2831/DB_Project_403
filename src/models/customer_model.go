package models

type CustomerType string

const (
	IndividualCustomerType  CustomerType = "individual"
	LegalEntityCustomerType CustomerType = "legal_entity"
)

type Customer struct {
	ID        int          `json:"id" sql:"id"`
	Type      CustomerType `json:"type" sql:"type"`
	ProfileID int          `json:"profile_id" sql:"profile_id"`
	Profile   *Profile     `json:"profile" sql:""`
	Account   []*Account   `json:"account" sql:"-"` // "" for relation , "-" not include
}

type CustomerWithFullName struct {
	FirstName string `json:"first_name" sql:"first_name"`
	LastName  string `json:"last_name" sql:"last_name"`
}

type CustomerWithFullNameAndAmount struct {
	FirstName    string  `json:"first_name" sql:"first_name"`
	LastName     string  `json:"last_name" sql:"last_name"`
	AccountType  string  `json:"account_type"  sql:"account_type"`
	TotalBalance float64 `json:"total_balance"  sql:"total_balance"`
}

type CustomerWithAccountNumber struct {
	FirstName     string `json:"first_name" sql:"first_name"`
	LastName      string `json:"last_name" sql:"last_name"`
	AccountNumber int    `json:"account_number" sql:"account_number"`
}

type CustomerWithMostLoan struct {
	ID         int    `json:"id" sql:"id"`
	FirstName  string `json:"first_name" sql:"first_name"`
	LastName   string `json:"last_name" sql:"last_name"`
	Type       string `json:"type" sql:"type"`
	LoanNumber int    `json:"loan_number" sql:"loan_number"`
}
