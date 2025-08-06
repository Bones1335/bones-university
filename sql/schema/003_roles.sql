-- +goose Up
CREATE TABLE roles (
  roles_id UUID PRIMARY KEY,
  role_name TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE roles;
