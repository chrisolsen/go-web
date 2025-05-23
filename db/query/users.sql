-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;
-- name: GetUsers :many
SELECT *
FROM users
ORDER BY name;
-- name: CreateUser :one
INSERT INTO users (name)
VALUES ($1)
RETURNING *;
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
