// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID         uuid.UUID
	Status     string
	Type       string
	ExpiryDate time.Time
	Amount     int32
	CreatedAt  sql.NullTime
	UpdatedAt  sql.NullTime
}

type Transaction struct {
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
	UpdatedAt                  sql.NullTime
}
