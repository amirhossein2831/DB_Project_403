DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'account_type') THEN
            CREATE TYPE account_type AS ENUM ('savings', 'current', 'business');
    END IF;
END $$;

DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'account_status') THEN
            CREATE TYPE account_status AS ENUM ('active', 'closed', 'pending');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS account (
         id SERIAL PRIMARY KEY,
         account_number VARCHAR(20) NOT NULL UNIQUE,
         type account_type NOT NULL,
         amount DECIMAL(15, 2) DEFAULT 0.00,
         status account_status NOT NULL DEFAULT 'pending',
         customer_id INT NOT NULL,
         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
         closed_at TIMESTAMP NULL,
         CONSTRAINT fk_customer FOREIGN KEY (customer_id) REFERENCES customer (id) ON DELETE CASCADE
);