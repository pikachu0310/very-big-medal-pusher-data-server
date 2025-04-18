-- name: CreateUser :exec
INSERT INTO users (username, password)
VALUES ($1, $2);

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = LAST_INSERT_ID();

-- name: AuthenticateUser :one
SELECT *
FROM users
WHERE username = $1
  AND password = $2
LIMIT 1;
