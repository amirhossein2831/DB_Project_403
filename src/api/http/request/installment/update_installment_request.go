package insallment

type UpdateInstallmentRequest struct {
	LoanID       *int     `json:"loan_id" validate:"omitempty,gt=0"`
	AmountPaid   *float64 `json:"amount_paid" validate:"omitempty,gt=0"`
	InterestPaid *float64 `json:"interest_paid" validate:"omitempty,gt=0"`
	TotalPaid    *float64 `json:"total_paid" validate:"omitempty,gt=0"`
	DueDate      *string  `json:"due_date" validate:"omitempty,datetime=2006-01-02"`
	PaidDate     *string  `json:"paid_date,omitempty" validate:"omitempty,datetime=2006-01-02"`
}
