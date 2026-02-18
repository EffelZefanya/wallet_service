-- Created in PSQL
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE wallets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    balance DECIMAL(15, 2) DEFAULT 0.00,
    version INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);



-- For this case, only WITHDRAWAL is available
CREATE TYPE tx_type AS ENUM ('WITHDRAWAL', 'DEPOSIT');
-- For this case, only SUCCESS and FAILED is available. PENDING usually is set if a confirmation is needed. Like waiting for a TOP UP through a payment gateway when depositing money.
CREATE TYPE tx_status AS ENUM ('PENDING', 'SUCCESS', 'FAILED');

CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    wallet_id UUID REFERENCES wallets(id),
    amount DECIMAL(15, 2) NOT NULL,
    transaction_type tx_type NOT NULL,
    status tx_status NOT NULL DEFAULT 'PENDING',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (id, username, email) VALUES 
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'john_doe', 'john@example.com'),
('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a33', 'jane_doe', 'jane@example.com');

INSERT INTO wallets (id, user_id, balance, version) VALUES
('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 1000.00, 1),
('b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a44', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a33', 5000.50, 1);