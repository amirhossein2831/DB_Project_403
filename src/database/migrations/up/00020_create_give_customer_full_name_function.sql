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
