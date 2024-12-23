package models

type CustomerType string

const (
	Individual  CustomerType = "individual"
	LegalEntity CustomerType = "legal_entity"
)

type Customer struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	ProfileID int    `json:"profile_id"`
}
