package models

type BankMembers struct {
	ID        int    `json:"id" sql:"id"`
	FirstName string `json:"first_name" sql:"first_name"`
	LastName  string `json:"last_name" sql:"last_name"`
	UserType  string `json:"user_type" sql:"user_type"`
	Phone     string `json:"phone" sql:"phone"`
	Email     string `json:"email" sql:"email"`
	Address   string `json:"address" sql:"address"`
}
