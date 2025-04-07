CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    api_key VARCHAR(255) NOT NULL UNIQUE,
    balance DECIMAL(10, 2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_accounts_api_key ON accounts (api_key);
CREATE INDEX idx_accounts_email ON accounts (email);

CREATE TABLE IF NOT EXISTS invoices (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id UUID NOT NULL REFERENCES accounts(id),
    amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    payment_method VARCHAR(255) NOT NULL,
    card_last_digits VARCHAR(4) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);