package account

import "DB_Project/src/models"

type UpdateAccountRequest struct {
	AccountNumber *string               `json:"account_number" validate:"omitempty"`
	Type          *models.AccountType   `json:"type" validate:"omitempty"`
	Amount        *float64              `json:"amount" validate:"omitempty"`
	Status        *models.AccountStatus `json:"status" validate:"omitempty"`
	CustomerID    *int                  `json:"customer_id" validate:"omitempty"`
}
