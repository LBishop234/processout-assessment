CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS transactions (
    id uuid PRIMARY KEY DEFAULT uuid_genereate_v4(),
    card_no VARCHAR(16) NOT NULL,
    expiry_month INT NOT NULL,
    expiry_year INT NOT NULL,
    cvv INT NOT NULL,
    trans_time TIMESTAMP NOT NULL,
    amount DECIMAL(10, 2) NOT NULL
);