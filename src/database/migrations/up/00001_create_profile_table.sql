CREATE TABLE  IF NOT EXISTS profile (
     id SERIAL PRIMARY KEY,
     first_name VARCHAR(100) NOT NULL,
     last_name VARCHAR(100) NOT NULL,
     birth_date DATE,
     phone VARCHAR(15),
     email VARCHAR(255) UNIQUE,
     address TEXT
);