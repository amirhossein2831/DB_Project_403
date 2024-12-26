package models

import "time"

type Profile struct {
	ID        int       `json:"id" sql:"id"`
	FirstName string    `json:"first_name" sql:"first_name"`
	LastName  string    `json:"last_name" sql:"last_name"`
	BirthDate time.Time `json:"birth_date" sql:"birth_date"`
	Phone     string    `json:"phone" sql:"phone"`
	Email     string    `json:"email" sql:"email"`
	Address   string    `json:"address" sql:"address"`
}
