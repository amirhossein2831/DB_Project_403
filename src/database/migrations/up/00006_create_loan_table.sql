DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'loan_type') THEN
            CREATE TYPE loan_type AS ENUM ('personal', 'mortgage', 'business');
    END IF;
END $$;

DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'loan_status') THEN
            CREATE TYPE loan_status AS ENUM ('pending', 'approved', 'repaid', 'defaulted');
    END IF;
END $$;

-- Create the loan table
CREATE TABLE IF NOT EXISTS loan (
        id SERIAL PRIMARY KEY,
        customer_id INT NOT NULL,
        type loan_type NOT NULL,
        status loan_status NOT NULL,
        amount DECIMAL(10, 2) NOT NULL,
        interest_rate FLOAT NOT NULL,
        repayment_period INT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        finished_at TIMESTAMP,
        CONSTRAINT fk_customer FOREIGN KEY (customer_id) REFERENCES customer(id) ON DELETE CASCADE
);
