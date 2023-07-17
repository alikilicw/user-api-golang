// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: user.queries.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (firstname,
   lastname,
   username,
   email,
   password,
   verification_code
    )
VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id, firstname, lastname, username, email, password, verification_code, verified, created_at
`

type CreateUserParams struct {
	Firstname        sql.NullString `json:"firstname"`
	Lastname         sql.NullString `json:"lastname"`
	Username         string         `json:"username"`
	Email            string         `json:"email"`
	Password         string         `json:"password"`
	VerificationCode sql.NullString `json:"verification_code"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Firstname,
		arg.Lastname,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.VerificationCode,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.VerificationCode,
		&i.Verified,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE from users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, firstname, lastname, username, email, password, verification_code, verified, created_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.VerificationCode,
		&i.Verified,
		&i.CreatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, firstname, lastname, username, email, password, verification_code, verified, created_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.VerificationCode,
		&i.Verified,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, firstname, lastname, username, email, password, verification_code, verified, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.VerificationCode,
		&i.Verified,
		&i.CreatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, firstname, lastname, username, email, password, verification_code, verified, created_at FROM users
ORDER BY id
LIMIT $1
OFFSET $2
`

type GetUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetUsers(ctx context.Context, arg GetUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Firstname,
			&i.Lastname,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.VerificationCode,
			&i.Verified,
			&i.CreatedAt,
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

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET username = $2, verified = $3, verification_code = $4
WHERE id = $1
RETURNING id, firstname, lastname, username, email, password, verification_code, verified, created_at
`

type UpdateUserParams struct {
	ID               int64          `json:"id"`
	Username         string         `json:"username"`
	Verified         sql.NullBool   `json:"verified"`
	VerificationCode sql.NullString `json:"verification_code"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Username,
		arg.Verified,
		arg.VerificationCode,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.VerificationCode,
		&i.Verified,
		&i.CreatedAt,
	)
	return i, err
}
