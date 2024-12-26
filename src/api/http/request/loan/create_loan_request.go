package loan

import (
	"DB_Project/src/models"
)

type CreateLoanRequest struct {
	CustomerID      int               `json:"customer_id" validate:"required,gt=0"`
	Type            models.LoanType   `json:"type" validate:"required,oneof=personal mortgage business"`
	Status          models.LoanStatus `json:"status" validate:"required,oneof=pending approved repaid defaulted"`
	Amount          float64           `json:"amount" validate:"required,gt=0"`
	InterestRate    float32           `json:"interest_rate" validate:"required,gt=0"`
	RepaymentPeriod int               `json:"repayment_period" validate:"required,gt=0"`
	FinishedAt      string            `json:"finished_at,omitempty" validate:"required,datetime=2006-01-02"`
}
