package customer

import (
	"DB_Project/src/models"
	"time"
)

type CreateCustomerRequest struct {
	FirstName    string              `json:"first_name"`
	LastName     string              `json:"last_name"`
	BirthDate    time.Time           `json:"birth_date"`
	Phone        string              `json:"phone"`
	Email        string              `json:"email"`
	Address      string              `json:"address"`
	CustomerType models.CustomerType `json:"customer_type"`
}
