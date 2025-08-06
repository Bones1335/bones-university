-- name: CreateRole :one
INSERT INTO roles (
  id,
  role_name
)
VALUES (
  gen_random_uuid(),
  $1
)
RETURNING *;
