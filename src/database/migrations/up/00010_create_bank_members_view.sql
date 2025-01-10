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
;
