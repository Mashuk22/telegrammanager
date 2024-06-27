-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;
-- name: ListUsers :many
SELECT *
FROM users
ORDER BY username;
-- name: CreateUser :execresult
INSERT INTO users (
        chat_id,
        username,
        first_name,
        last_name,
        role_id,
        is_subscribed
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: UpdateUser :exec
UPDATE users
set chat_id = $2,
    username = $3,
    first_name = $4,
    last_name = $5,
    role_id = $6,
    is_subscribed = $7
WHERE id = $1;
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;