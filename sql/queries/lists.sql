-- name: GetMovieListByUser :one
SELECT list from movie_lists
WHERE user_id = $1;


-- name: GetCompletedListByUser :one
SELECT list from completed_titles
WHERE user_id = $1;

-- name: UpdateMovieList :exec
UPDATE movie_lists
  set list = $1
WHERE user_id = $2
RETURNING *;

-- name: UpdateCompletedList :exec
UPDATE completed_titles
  set list = list || $1 ::jsonb
WHERE user_id = $2
RETURNING *;

-- name: GetGameListByUser :one
SELECT list from game_lists
WHERE user_id = $1;


-- name: UpdateGameList :exec
UPDATE game_lists
  set list = $1
WHERE user_id = $2
RETURNING *;

-- name: GetBookListByUser :one
SELECT list from book_lists
WHERE user_id = $1;

-- name: UpdateBookList :exec
UPDATE book_lists
  set list = $1
WHERE user_id = $2
RETURNING *;

-- name: GetShowListByUser :one
SELECT list from show_lists
WHERE user_id = $1;

-- name: UpdateShowList :exec
UPDATE show_lists
  set list = $1
WHERE user_id = $2
RETURNING *;
