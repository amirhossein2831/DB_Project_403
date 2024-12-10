DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'customer_type') THEN
            CREATE TYPE customer_type AS ENUM ('individual', 'legal_entity');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS customer(
       id SERIAL PRIMARY KEY,
       type customer_type NOT NULL,
       profile_id INT NOT NULL,
       CONSTRAINT fk_profile FOREIGN KEY (profile_id) REFERENCES profile (id) ON DELETE CASCADE
);
