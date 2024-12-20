CREATE TABLE IF NOT EXISTS employee(
      id SERIAL PRIMARY KEY,
      position VARCHAR(100) NOT NULL,
      profile_id INT UNIQUE NOT NULL,
      CONSTRAINT fk_profile FOREIGN KEY (profile_id) REFERENCES profile (id) ON DELETE CASCADE
);
