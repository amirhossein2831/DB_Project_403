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
