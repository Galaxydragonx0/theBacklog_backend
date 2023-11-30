-- +goose Up

CREATE TABLE users(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    email VARCHAR(64) NOT NULL UNIQUE,
    api_key VARCHAR(64) UNIQUE NOT NULL Default (encode(sha256(random()::text::bytea), 'hex'))
);



-- +goose Down
DROP TABLE users;