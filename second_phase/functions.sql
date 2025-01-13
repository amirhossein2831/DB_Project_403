CREATE OR REPLACE FUNCTION prevent_delete_with_active_loans()
RETURNS TRIGGER AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM loan WHERE customer_id = OLD.id AND status <> 'repaid') THEN
        RAISE EXCEPTION 'Cannot delete customer with active loans';
END IF;

RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_prevent_delete_customer
    BEFORE DELETE ON customer
    FOR EACH ROW
    EXECUTE FUNCTION prevent_delete_with_active_loans();


CREATE OR REPLACE FUNCTION update_account_balance_after_transaction()
RETURNS TRIGGER AS $$
BEGIN

    IF NEW.type = 'transfer' AND NEW.destination_account_id IS NULL THEN
        RAISE EXCEPTION 'Transfer type transaction must have a valid destination account'
        USING ERRCODE = 'P0002';
END IF;

    IF NEW.type = 'withdrawal' OR NEW.type = 'transfer' THEN
UPDATE account
SET amount = amount - NEW.amount
WHERE id = NEW.source_account_id;
END IF;

    IF NEW.type = 'deposit' OR NEW.type = 'transfer' THEN
UPDATE account
SET amount = amount + NEW.amount
WHERE id = NEW.destination_account_id;
END IF;

RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trg_update_account_balance_after_transaction
    AFTER INSERT ON transaction
    FOR EACH ROW
    EXECUTE FUNCTION update_account_balance_after_transaction();


CREATE OR REPLACE FUNCTION check_sufficient_funds()
    RETURNS TRIGGER AS $$
BEGIN
    IF NEW.type = 'withdrawal' OR NEW.type = 'transfer' THEN
        IF (SELECT amount FROM account WHERE id = NEW.source_account_id) < NEW.amount THEN
            RAISE EXCEPTION 'Insufficient funds in the source account'
            USING ERRCODE = 'P0001';
END IF;
END IF;

RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_check_sufficient_funds
    BEFORE INSERT ON transaction
    FOR EACH ROW
    EXECUTE FUNCTION check_sufficient_funds();


CREATE OR REPLACE FUNCTION get_total_amount(customer_id INT)
RETURNS DECIMAL(15, 2) AS $$
    DECLARE
total_amount DECIMAL(15, 2) := 0.00;
BEGIN
SELECT SUM(amount) INTO total_amount
FROM account
WHERE get_total_amount.customer_id = account.customer_id AND status = 'active';

RETURN total_amount;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION get_loan_status(loan_id INT)
RETURNS TEXT AS $$
    DECLARE
loan_status TEXT;
BEGIN
SELECT
    CASE
        WHEN finished_at < CURRENT_DATE THEN 'repaid'
        ELSE status
        END
INTO loan_status
FROM loan
WHERE id = loan_id;

IF loan_status IS NULL THEN
                RAISE EXCEPTION 'Loan with ID % not found', loan_id;
END IF;

RETURN loan_status;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION get_active_loans_count(customer_id INT)
RETURNS INT AS $$
    DECLARE
active_loans_count INT;
BEGIN
SELECT COUNT(*) INTO active_loans_count
FROM loan
WHERE loan.customer_id = get_active_loans_count.customer_id
  AND (status != 'repaid' AND status != 'defaulted')
  AND finished_at IS NULL;

RETURN active_loans_count;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION get_total_payments(loan_id INT)
RETURNS DECIMAL(10, 2) AS $$
    DECLARE
total_payments DECIMAL(10, 2) := 0.00;
BEGIN
SELECT SUM(total_paid) INTO total_payments
FROM installment
WHERE get_total_payments.loan_id = installment.loan_id AND installment.paid_date IS NOT NULL;

RETURN total_payments;
END;
$$ LANGUAGE plpgsql;
