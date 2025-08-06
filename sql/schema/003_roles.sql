-- +goose Up
CREATE TABLE roles (
  id UUID PRIMARY KEY,
  role_name TEXT UNIQUE NOT NULL
);

-- +goose Down
DELETE TABLE roles;
