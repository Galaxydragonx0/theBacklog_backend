-- name: GetUserByApiKey :one
SELECT * from users WHERE api_key = $1;


-- name: CreateUser :one
INSERT into users(id, created_at, updated_at, email, api_key)
VALUES ($1, $2, $3, $4, encode (sha256(random()::text::bytea), 'hex'))
RETURNING *;

