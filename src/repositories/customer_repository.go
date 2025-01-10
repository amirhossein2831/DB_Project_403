package repositories

import (
	"DB_Project/src/database/connection/pgx"
	"DB_Project/src/models"
	"DB_Project/src/utils"
	"context"
	"fmt"
)

type CustomerRepository struct {
	AccountRepository *AccountRepository
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{
		AccountRepository: NewAccountRepository(),
	}
}

func (repository *CustomerRepository) List() ([]*models.Customer, error) {
	var customers []*models.Customer
	rows, err := pgx.GetInstance().Query(context.Background(), "SELECT * FROM customer c INNER JOIN profile p ON c.profile_id = p.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.Customer
		err = utils.FillStructFromRowsWithJoin(rows, &customer)
		if err != nil {
			return nil, err
		}
		customers = append(customers, &customer)
	}

	for _, customer := range customers {
		var accounts []*models.Account
		accounts, err = repository.AccountRepository.ListByCustomerId(customer.ID)
		if err != nil {
			return nil, err
		}

		customer.Account = accounts
	}

	return customers, rows.Err()
}

func (repository *CustomerRepository) ListWithFullName() ([]*models.CustomerWithFullName, error) {
	var customers []*models.CustomerWithFullName
	rows, err := pgx.GetInstance().Query(context.Background(), "SELECT p.first_name,p.last_name FROM customer c INNER JOIN profile p ON c.profile_id = p.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.CustomerWithFullName
		err = utils.FillStructFromRows(rows, &customer)

		customers = append(customers, &customer)
	}

	return customers, nil
}

func (repository *CustomerRepository) ListWithFullNameAndTotalAmount() ([]*models.CustomerWithFullNameAndAmount, error) {
	var customers []*models.CustomerWithFullNameAndAmount
	rows, err := pgx.GetInstance().Query(context.Background(), `
	SELECT
		p.first_name AS first_name,
		p.last_name AS last_name,
		a.type AS account_type,
		SUM(a.amount) AS total_balance
	FROM profile p
		INNER JOIN customer c ON p.id = c.profile_id
		INNER JOIN account a ON c.id = a.customer_id
	GROUP BY p.first_name, p.last_name, a.type
	ORDER BY p.first_name, p.last_name, a.type;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.CustomerWithFullNameAndAmount
		err = utils.FillStructFromRows(rows, &customer)

		customers = append(customers, &customer)
	}

	return customers, nil
}

func (repository *CustomerRepository) ListWithFullNameAndAccountNumber() ([]*models.CustomerWithAccountNumber, error) {
	var customers []*models.CustomerWithAccountNumber
	rows, err := pgx.GetInstance().Query(context.Background(), `
	SELECT p.first_name,p.last_name,count(a.id) as account_number 
	From customer c 
	    INNER JOIN public.profile p on c.profile_id = p.id 
	    INNER JOIN public.account a on c.id = a.customer_id 
	GROUP BY c.id,p.first_name, p.last_name 
	HAVING count(a.id) > 1;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.CustomerWithAccountNumber
		err = utils.FillStructFromRows(rows, &customer)

		customers = append(customers, &customer)
	}

	return customers, nil
}

func (repository *CustomerRepository) ListWithMostLoan() ([]*models.CustomerWithMostLoan, error) {
	var customers []*models.CustomerWithMostLoan
	rows, err := pgx.GetInstance().Query(context.Background(), `
	WITH LoanCounts AS (
		SELECT
			c.id AS id,
			c.type AS type,
			COUNT(l.id) AS loan_count
		FROM
			customer c
				LEFT JOIN loan l ON c.id = l.customer_id
		GROUP BY
			c.id, c.type
	),
		 MaxLoanCount AS (
			 SELECT
				 MAX(loan_count) AS max_loan_count
			 FROM
				 LoanCounts
		 )
	SELECT
		lc.id,
		p.first_name,
		p.last_name,
		lc.type,
		lc.loan_count
	FROM
		LoanCounts lc
			CROSS JOIN MaxLoanCount mlc
			INNER JOIN profile p on p.id = lc.id
	WHERE
		lc.loan_count = mlc.max_loan_count;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.CustomerWithMostLoan
		err = utils.FillStructFromRows(rows, &customer)

		customers = append(customers, &customer)
	}

	return customers, nil
}

func (repository *CustomerRepository) ListWithInstallmentsPenalty() ([]*models.CustomerWithPenaltyInstallment, error) {
	var customers []*models.CustomerWithPenaltyInstallment
	rows, err := pgx.GetInstance().Query(context.Background(), `
	SELECT
		p.first_name,
		p.last_name,
		l.id AS loan_id,
		l.amount AS loan_amount
	FROM
		customer c
			JOIN
		profile p ON c.profile_id = p.id
			JOIN
		loan l ON c.id = l.customer_id
			JOIN
		installment i ON l.id = i.loan_id
	WHERE
		(i.paid_date IS NULL AND i.due_date < CURRENT_DATE)
	   OR
		(i.paid_date > i.due_date)
	group by l.id, p.last_name, p.first_name, l.amount;
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.CustomerWithPenaltyInstallment
		err = utils.FillStructFromRows(rows, &customer)

		customers = append(customers, &customer)
	}

	return customers, nil
}

func (repository *CustomerRepository) ListWithMostAmount() ([]*models.CustomerWithMostAmount, error) {
	var customers []*models.CustomerWithMostAmount
	rows, err := pgx.GetInstance().Query(context.Background(), `
	SELECT
		p.first_name,
		p.last_name,
		SUM(a.amount) AS total_amount
	FROM
		customer c
			JOIN
		profile p ON c.profile_id = p.id
			JOIN
		account a ON c.id = a.customer_id
	GROUP BY
		p.first_name, p.last_name
	ORDER BY
		total_amount DESC
	LIMIT 5;
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.CustomerWithMostAmount
		err = utils.FillStructFromRows(rows, &customer)

		customers = append(customers, &customer)
	}

	return customers, nil
}

func (repository *CustomerRepository) Get(id string) (*models.Customer, error) {
	var customer models.Customer
	row := pgx.GetInstance().QueryRow(context.Background(), "SELECT * FROM customer c LEFT JOIN profile p ON c.profile_id = p.id WHERE c.id=$1", id)
	err := utils.FillStructFromRowWithJoin(row, &customer)
	if err != nil {
		return nil, err
	}

	var accounts []*models.Account
	accounts, err = repository.AccountRepository.ListByCustomerId(customer.ID)
	if err != nil {
		return nil, err
	}
	customer.Account = accounts

	return &customer, nil
}

func (repository *CustomerRepository) Create(customer *models.Customer, profile *models.Profile) error {
	// context
	ctx := context.Background()

	// Start a transaction
	tx, err := pgx.GetInstance().Begin(ctx)
	if err != nil {
		return err
	}
	// Ensure that the transaction is rolled back if any error occurs
	defer tx.Rollback(ctx)

	// Insert profile and retrieve the generated profile ID
	err = tx.QueryRow(ctx, "INSERT INTO profile (first_name, last_name, birth_date, phone, email, address) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		profile.FirstName, profile.LastName, profile.BirthDate, profile.Phone, profile.Email, profile.Address,
	).Scan(&profile.ID)
	if err != nil {
		return err
	}

	// Insert customer using the generated profile ID
	_, err = tx.Exec(ctx, "INSERT INTO customer (profile_id, type) VALUES ($1, $2)",
		profile.ID, customer.Type,
	)
	if err != nil {
		return err
	}

	// Commit the transaction if everything is successful
	return tx.Commit(ctx)
}

func (repository *CustomerRepository) UpdateField(name, id string, value interface{}) error {
	_, err := pgx.GetInstance().Exec(context.Background(), fmt.Sprintf("UPDATE customer SET %s = $1 WHERE id = $2", name), value, id)
	return err
}

func (repository *CustomerRepository) Delete(id string) error {
	_, err := pgx.GetInstance().Exec(context.Background(), "DELETE FROM customer WHERE id=$1", id)
	return err
}
