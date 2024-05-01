-- name: GetUserByApiKey :one
SELECT * from users WHERE api_key = $1;


-- name: CreateUser :one
INSERT into users(id, created_at, updated_at, email, api_key)
VALUES ($1, $2, $3, $4, encode (sha256(random()::text::bytea), 'hex'))
RETURNING *;

-- name: CreateMovieList :exec
Insert into movie_lists(id, user_id, list)
VALUES ($1, $2, $3);

-- name: CreateShowList :exec
Insert into show_lists(id, user_id, list)
VALUES ($1, $2, $3);

-- name: CreateBookList :exec
Insert into book_lists(id, user_id, list)
VALUES ($1, $2, $3);

-- name: CreateGameList :exec
Insert into game_lists(id, user_id, list)
VALUES ($1, $2, $3);

-- name: GetUser :one
SELECT * FROM users
WHERE email =$1;

-- name: CreateCompletedList :exec
Insert into completed_titles(id, user_id, list)
VALUES ($1, $2, $3);