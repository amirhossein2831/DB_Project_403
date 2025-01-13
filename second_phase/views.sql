CREATE VIEW customer_accounts AS
SELECT
    p.first_name,
    p.last_name,
    p.phone,
    a.account_number,
    a.type,
    a.amount
FROM
    customer c
        INNER JOIN
    profile p ON c.profile_id = p.id
        INNER JOIN
    account a ON c.id = a.customer_id;


CREATE VIEW bank_transactions AS
SELECT t.id,
       t.type,
       sa.account_number as source_account_number,
       da.account_number as destination_account_number,
       t.amount,
       t.created_at
FROM transaction t
         INNER JOIN public.account sa ON t.source_account_id = sa.id
         LEFT JOIN public.account da ON t.destination_account_id = da.id
ORDER BY t.id
;


CREATE VIEW bank_members AS
SELECT
    p.id,
    p.first_name,
    p.last_name,
    CASE
        WHEN e.id IS NOT NULL THEN 'employee'
        WHEN c.id IS NOT NULL THEN 'customer'
        END AS user_type,
    p.phone,
    p.email,
    p.address
FROM
    profile p
        LEFT JOIN employee e ON p.id = e.profile_id
        LEFT JOIN customer c ON p.id = c.profile_id
ORDER BY p.id
;


CREATE OR REPLACE FUNCTION set_account_creation_date()
RETURNS TRIGGER AS $$
BEGIN
    NEW.DateOpened := CURRENT_TIMESTAMP;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trg_set_account_creation_date
    BEFORE INSERT ON Account
    FOR EACH ROW
EXECUTE FUNCTION set_account_creation_date();


CREATE OR REPLACE FUNCTION get_customer_name_by_id(customer_id INT)
RETURNS TABLE(first_name VARCHAR, last_name VARCHAR) AS $$
BEGIN
RETURN QUERY
SELECT p.first_name, p.last_name
FROM profile p
         JOIN customer c ON p.id = c.profile_id
WHERE c.id = customer_id;
END;
$$ LANGUAGE plpgsql;

