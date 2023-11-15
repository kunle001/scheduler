-- name: CreateTransactionJob :one
INSERT INTO transactions (
  id, 
  status, 
  originator_account_number, 
  beneficiary_account_name, 
  beneficiary_account_number, 
  destination_institution_name,
  destination_institution_code,
  customer_account_name,
  narration,
  amount,
  end_date,
  created_at
)

VALUES ($1, $2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
RETURNING *;

-- name: FindValidtransactions :many
SELECT * FROM transactions 
WHERE end_date > CURRENT_DATE AND status = 'ongoing';