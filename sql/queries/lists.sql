-- name: GetMovieListByUser :one
SELECT list from movie_lists
WHERE user_id = $1;


-- name: UpdateMovieList :exec
UPDATE movie_lists
  set list = $1
WHERE user_id = $2
RETURNING *;
