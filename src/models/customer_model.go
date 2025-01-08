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
