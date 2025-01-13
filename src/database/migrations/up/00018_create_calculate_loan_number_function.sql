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
