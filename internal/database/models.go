// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Task struct {
	ID         uuid.UUID
	Status     sql.NullString
	Type       sql.NullString
	ExpiryDate sql.NullTime
	CreatedAt  sql.NullTime
	UpdatedAt  sql.NullTime
}
