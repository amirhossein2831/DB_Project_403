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
;
