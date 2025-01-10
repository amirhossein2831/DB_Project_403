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
