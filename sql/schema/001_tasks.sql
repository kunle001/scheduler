-- +goose Up
CREATE TABLE tasks (
  id UUID PRIMARY KEY,
  status TEXT CHECK (status IN ('fail', 'success', 'pending')),
  type TEXT CHECK (type IN ('bill', 'transfer', 'airtime', 'data')),
  expiry_date TIMESTAMP,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE tasks;
