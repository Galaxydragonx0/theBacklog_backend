-- +goose Up

CREATE TABLE movie_lists(
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    list JSONB NOT NULL DEFAULT '[{}]'::JSONB
);


CREATE TABLE show_lists(
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    list JSONB NOT NULL DEFAULT '[{}]'::JSONB
);


CREATE TABLE book_lists(
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    list JSONB NOT NULL DEFAULT '[{}]'::JSONB
);


CREATE TABLE game_lists(
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    list JSONB NOT NULL DEFAULT '[{}]'::JSONB
);




-- +goose Down
DROP TABLE movie_lists;
DROP TABLE show_lists;
DROP TABLE book_lists;
DROP TABLE game_lists;