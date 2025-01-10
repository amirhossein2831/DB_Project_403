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
