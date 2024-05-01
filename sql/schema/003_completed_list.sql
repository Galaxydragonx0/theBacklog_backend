-- +goose Up
CREATE TABLE completed_titles(
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    list JSONB NOT NULL DEFAULT '[]' ::JSONB
);



-- +goose Down
DROP TABLE completed_titles
