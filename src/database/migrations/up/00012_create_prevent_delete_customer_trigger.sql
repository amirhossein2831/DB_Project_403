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
