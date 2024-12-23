package models

type CustomerType string

const (
	Individual  CustomerType = "individual"
	LegalEntity CustomerType = "legal_entity"
)

type Customer struct {
	ID        int      `json:"id" sql:"id"`
	Type      string   `json:"type" sql:"type"`
	ProfileID int      `json:"profile_id" sql:"profile_id"`
	Profile   *Profile `json:"profile" sql:""`
}
