// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package dbx

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type AuthUser struct {
	ID             pgtype.UUID
	Email          string
	HashedPassword []byte
	FirstName      string
	LastName       string
	IsActive       bool
	IsVerified     bool
	CreatedAt      pgtype.Timestamptz
	UpdatedAt      pgtype.Timestamptz
}
