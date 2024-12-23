package models

type Profile struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate string `json:"birth_date"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
}
