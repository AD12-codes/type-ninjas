// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: users.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const getAllUsers = `-- name: GetAllUsers :many
SELECT
  id, auth0_id, email, username, first_name, last_name, is_deleted, deleted_at, created_at, updated_at
FROM
  users
WHERE
  is_deleted = FALSE
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Auth0ID,
			&i.Email,
			&i.Username,
			&i.FirstName,
			&i.LastName,
			&i.IsDeleted,
			&i.DeletedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserById = `-- name: GetUserById :one
SELECT
  id, auth0_id, email, username, first_name, last_name, is_deleted, deleted_at, created_at, updated_at
FROM
  users
WHERE
  id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Auth0ID,
		&i.Email,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.IsDeleted,
		&i.DeletedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const registerUser = `-- name: RegisterUser :exec
INSERT INTO users(id, auth0_id, email, username, first_name, last_name)
  VALUES ($1, $2, $3, $4, $5, $6)
`

type RegisterUserParams struct {
	ID        uuid.UUID `json:"id"`
	Auth0ID   string    `json:"auth0Id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}

func (q *Queries) RegisterUser(ctx context.Context, arg RegisterUserParams) error {
	_, err := q.db.Exec(ctx, registerUser,
		arg.ID,
		arg.Auth0ID,
		arg.Email,
		arg.Username,
		arg.FirstName,
		arg.LastName,
	)
	return err
}
