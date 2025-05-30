-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    last_name TEXT NOT NULL,
    first_name TEXT NOT NULL,
    username TEXT UNIQUE NOT NULL,
    personal_email TEXT UNIQUE NOT NULL,
    university_email TEXT UNIQUE NOT NULL,
    isAdmin BOOLEAN NOT NULL,
    password TEXT NOT NULL DEFAULT 'unset'
);

-- +goose Down
DROP TABLE users;