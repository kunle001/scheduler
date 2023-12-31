// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: transactions.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createTransactionJob = `-- name: CreateTransactionJob :one
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
RETURNING id, status, originator_account_number, beneficiary_account_name, beneficiary_account_number, destination_institution_name, destination_institution_code, customer_account_name, narration, amount, end_date, created_at, updated_at
`

type CreateTransactionJobParams struct {
	ID                         uuid.UUID
	Status                     string
	OriginatorAccountNumber    string
	BeneficiaryAccountName     string
	BeneficiaryAccountNumber   string
	DestinationInstitutionName string
	DestinationInstitutionCode string
	CustomerAccountName        string
	Narration                  string
	Amount                     int32
	EndDate                    time.Time
	CreatedAt                  time.Time
}

func (q *Queries) CreateTransactionJob(ctx context.Context, arg CreateTransactionJobParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, createTransactionJob,
		arg.ID,
		arg.Status,
		arg.OriginatorAccountNumber,
		arg.BeneficiaryAccountName,
		arg.BeneficiaryAccountNumber,
		arg.DestinationInstitutionName,
		arg.DestinationInstitutionCode,
		arg.CustomerAccountName,
		arg.Narration,
		arg.Amount,
		arg.EndDate,
		arg.CreatedAt,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.OriginatorAccountNumber,
		&i.BeneficiaryAccountName,
		&i.BeneficiaryAccountNumber,
		&i.DestinationInstitutionName,
		&i.DestinationInstitutionCode,
		&i.CustomerAccountName,
		&i.Narration,
		&i.Amount,
		&i.EndDate,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findValidtransactions = `-- name: FindValidtransactions :many
SELECT id, status, originator_account_number, beneficiary_account_name, beneficiary_account_number, destination_institution_name, destination_institution_code, customer_account_name, narration, amount, end_date, created_at, updated_at FROM transactions 
WHERE end_date > CURRENT_DATE AND status = 'ongoing'
`

func (q *Queries) FindValidtransactions(ctx context.Context) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, findValidtransactions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.Status,
			&i.OriginatorAccountNumber,
			&i.BeneficiaryAccountName,
			&i.BeneficiaryAccountNumber,
			&i.DestinationInstitutionName,
			&i.DestinationInstitutionCode,
			&i.CustomerAccountName,
			&i.Narration,
			&i.Amount,
			&i.EndDate,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
