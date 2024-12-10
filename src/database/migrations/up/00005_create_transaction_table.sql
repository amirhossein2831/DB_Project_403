DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_type') THEN
            CREATE TYPE transaction_type AS ENUM ('withdrawal', 'deposit', 'transfer');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS transaction (
       id SERIAL PRIMARY KEY,
       type transaction_type NOT NULL,
       amount DECIMAL(10, 2) NOT NULL,
       source_account_id INT NOT NULL,
       destination_account_id INT,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       CONSTRAINT fk_source_account FOREIGN KEY (source_account_id) REFERENCES account(id) ON DELETE CASCADE,
       CONSTRAINT fk_destination_account FOREIGN KEY (destination_account_id) REFERENCES account(id) ON DELETE CASCADE
);
