-- +goose Up
CREATE TABLE transactions (
  id UUID PRIMARY KEY,
  status TEXT NOT NULL CHECK (status IN ('ongoing', 'completed')),
  originator_account_number  TEXT NOT NULL,
  beneficiary_account_name  TEXT NOT NULL,
  beneficiary_account_number  TEXT NOT NULL,
  destination_institution_name  TEXT NOT NULL,
  destination_institution_code  TEXT NOT NULL,
  customer_account_name  TEXT NOT NULL,
  narration  TEXT NOT NULL,
  amount INTEGER NOT NULL,
  end_date TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE transactions;