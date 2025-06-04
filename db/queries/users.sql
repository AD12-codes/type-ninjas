-- name: RegisterUser :exec
INSERT INTO users(id, auth0_id, email, username, first_name, last_name)
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
  users
WHERE
  is_deleted = FALSE;

