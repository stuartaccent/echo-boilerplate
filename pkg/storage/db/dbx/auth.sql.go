// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: auth.sql

package dbx

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO auth_users (email, hashed_password, first_name, last_name)
VALUES ($1, $2, $3, $4)
RETURNING id, email, hashed_password, first_name, last_name, is_active, is_verified, created_at, updated_at
`

type CreateUserParams struct {
	Email          string
	HashedPassword []byte
	FirstName      string
	LastName       string
}

// create a new user
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (AuthUser, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Email,
		arg.HashedPassword,
		arg.FirstName,
		arg.LastName,
	)
	var i AuthUser
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.IsActive,
		&i.IsVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, hashed_password, first_name, last_name, is_active, is_verified, created_at, updated_at
FROM auth_users
WHERE email = $1
LIMIT 1
`

// get a user by their email
func (q *Queries) GetUserByEmail(ctx context.Context, email string) (AuthUser, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i AuthUser
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.IsActive,
		&i.IsVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, email, hashed_password, first_name, last_name, is_active, is_verified, created_at, updated_at
FROM auth_users
WHERE id = $1
LIMIT 1
`

// get a user by their id
func (q *Queries) GetUserByID(ctx context.Context, id pgtype.UUID) (AuthUser, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i AuthUser
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.IsActive,
		&i.IsVerified,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const setUserPasswordByEmail = `-- name: SetUserPasswordByEmail :exec
UPDATE auth_users
SET hashed_password = $2
WHERE email = $1
`

type SetUserPasswordByEmailParams struct {
	Email          string
	HashedPassword []byte
}

// set a user's password
func (q *Queries) SetUserPasswordByEmail(ctx context.Context, arg SetUserPasswordByEmailParams) error {
	_, err := q.db.Exec(ctx, setUserPasswordByEmail, arg.Email, arg.HashedPassword)
	return err
}
