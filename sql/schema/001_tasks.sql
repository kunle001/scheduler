-- +goose Up
CREATE TABLE tasks (
  id UUID PRIMARY KEY,
  status TEXT NOT NULL CHECK (status IN ('ongoing', 'completed')),
  type TEXT NOT NULL CHECK (type IN ('bill', 'transfer', 'airtime', 'data')),
  expiry_date TIMESTAMP NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE tasks;
