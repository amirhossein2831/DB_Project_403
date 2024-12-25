package models

import "github.com/jackc/pgx/v4"

type CustomerType string

const (
	IndividualCustomerType  CustomerType = "individual"
	LegalEntityCustomerType CustomerType = "legal_entity"
)

type Customer struct {
	ID        int          `json:"id" sql:"id"`
	Type      CustomerType `json:"type" sql:"type"`
	ProfileID int          `json:"profile_id" sql:"profile_id"`
	Profile   *Profile     `json:"profile" sql:""`
	Account   []*Account   `json:"account" sql:""`
}

func FetchCustomersAccount(customers []*Customer) ([]*Customer, error) {
	customersMap := make(map[int]*Customer)
	customersSlice := make([]*Customer, 0)

	if len(customers) == 0 {
		return nil, nil
	}

	for _, customer := range customers {
		if existingCustomer, ok := customersMap[customer.ID]; ok {
			existingCustomer.Account = append(existingCustomer.Account, customer.Account...)
		} else {
			customersMap[customer.ID] = customer
			customersSlice = append(customersSlice, customer)
		}
	}

	return customersSlice, nil
}

func FetchCustomerAccount(customers []*Customer) (*Customer, error) {
	if len(customers) == 0 {
		return nil, pgx.ErrNoRows
	}

	if len(customers) == 1 {
		return customers[0], nil
	}

	customerStruct := customers[0]
	for i := 1; i < len(customers); i++ {
		customerStruct.Account = append(customerStruct.Account, customers[i].Account...)
	}

	return customerStruct, nil
}
