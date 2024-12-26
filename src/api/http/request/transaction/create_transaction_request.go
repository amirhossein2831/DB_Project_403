package transaction

type CreateTransactionRequest struct {
	Type                 string  `json:"type" validate:"required,oneof=deposit withdrawal transfer"`
	Amount               float64 `json:"amount" validate:"required,gt=0"`
	SourceAccountId      int     `json:"source_account_id" validate:"required,gt=0"`
	DestinationAccountId int     `json:"destination_account_id" validate:"omitempty,gt=0"`
}
