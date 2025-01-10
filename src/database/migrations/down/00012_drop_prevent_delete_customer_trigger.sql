DROP TRIGGER IF EXISTS trg_prevent_delete_customer ON customer;

DROP FUNCTION IF EXISTS prevent_delete_with_active_loans();
