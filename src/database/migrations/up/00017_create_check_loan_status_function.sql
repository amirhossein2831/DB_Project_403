CREATE OR REPLACE FUNCTION get_loan_status(loan_id INT)
RETURNS TEXT AS $$
    DECLARE
    loan_status TEXT;
    BEGIN
        SELECT
            CASE
                WHEN finished_at IS NOT NULL THEN 'repaid'
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