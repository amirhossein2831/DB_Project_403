CREATE TABLE IF NOT EXISTS installment (
       id SERIAL PRIMARY KEY,
       loan_id INT NOT NULL,
       amount_paid DECIMAL(10, 2) NOT NULL,
       interest_paid DECIMAL(10, 2) NOT NULL,
       total_paid DECIMAL(10, 2) NOT NULL,
       due_date DATE NOT NULL,
       paid_date DATE,
       CONSTRAINT fk_loan FOREIGN KEY (loan_id) REFERENCES loan(id) ON DELETE CASCADE
);  