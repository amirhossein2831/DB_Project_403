package transaction

type UpdateTransactionRequest struct {
	Type                 *string  `json:"type" validate:"omitempty,oneof='' 'deposit' 'withdrawal' 'transfer'"`
	Amount               *float64 `json:"amount" validate:"omitempty,gt=0"`
	SourceAccountId      *int     `json:"source_account_id" validate:"omitempty,gt=0"`
	DestinationAccountId *int     `json:"destination_account_id" validate:"omitempty,gt=0"`
}
