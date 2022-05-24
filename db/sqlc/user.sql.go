// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
username,
hashed_password,
first_name,
last_name,
email,
contact_number,
address
) VALUES (
$1, $2, $3, $4, $5, $6, $7
)
RETURNING username, hashed_password, first_name, last_name, email, contact_number, address, updated_at, created_at
`

type CreateUserParams struct {
	Username       string         `json:"username"`
	HashedPassword string         `json:"hashed_password"`
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	Email          string         `json:"email"`
	ContactNumber  sql.NullString `json:"contact_number"`
	Address        sql.NullString `json:"address"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.ContactNumber,
		arg.Address,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.ContactNumber,
		&i.Address,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT username, hashed_password, first_name, last_name, email, contact_number, address, updated_at, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.ContactNumber,
		&i.Address,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}