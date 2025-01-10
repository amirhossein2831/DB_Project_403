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
