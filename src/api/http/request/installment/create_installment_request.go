package insallment

type CreateInstallmentRequest struct {
	LoanID       int     `json:"loan_id" validate:"required,gt=0"`
	AmountPaid   float64 `json:"amount_paid" validate:"required,gt=0"`
	InterestPaid float64 `json:"interest_paid" validate:"required,gt=0"`
	TotalPaid    float64 `json:"total_paid" validate:"required,gt=0"`
	DueDate      string  `json:"due_date" validate:"required,datetime=2006-01-02"`
	PaidDate     *string `json:"paid_date,omitempty" validate:"omitempty,datetime=2006-01-02"`
}
