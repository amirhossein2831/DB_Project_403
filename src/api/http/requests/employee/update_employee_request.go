package employee

type UpdateEmployeeRequest struct {
	FirstName *string `json:"first_name" validate:"omitempty"`
	LastName  *string `json:"last_name" validate:"omitempty"`
	BirthDate *string `json:"birth_date" validate:"omitempty,datetime=2006-01-02"`
	Phone     *string `json:"phone" validate:"omitempty"`
	Email     *string `json:"email" validate:"omitempty,email"`
	Address   *string `json:"address" validate:"omitempty"`
	Position  *string `json:"position" validate:"omitempty"`
}
