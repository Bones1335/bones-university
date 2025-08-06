-- +goose Up
CREATE TABLE users_roles (
  users_roles_id UUID PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES users (users_id) ON DELETE CASCADE,
  role_id UUID NOT NULL REFERENCES roles (roles_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE users_roles;
