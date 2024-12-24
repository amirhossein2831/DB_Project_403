package account

import (
	"DB_Project/src/models"
)

type CreateAccountRequest struct {
	AccountNumber string               `json:"account_number" validate:"required"`
	Type          models.AccountType   `json:"type" validate:"required"`
	Amount        float64              `json:"amount" validate:"required,gte=0"`
	Status        models.AccountStatus `json:"status" validate:"required"`
	CustomerID    int                  `json:"customer_id" validate:"required,gt=0"`
}
