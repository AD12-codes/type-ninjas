-- name: CreateUser :exec
INSERT INTO users(id, first_name, last_name, email, username, password)
  VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetUserById :one
SELECT
  *
FROM
  users
WHERE
  id = $1;

-- name: GetAllUsers :many
SELECT
  *
FROM
  users;

-- name: DeleteUser :exec
UPDATE
  users
SET
  deleted_at = $2,
  is_deleted = $3
WHERE
  id = $1;

