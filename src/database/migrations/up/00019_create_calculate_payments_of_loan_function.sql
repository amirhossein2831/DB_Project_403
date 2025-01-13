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
