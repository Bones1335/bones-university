-- name: CreateUsersRoles :one 
INSERT INTO users_roles (
  users_roles_id,
  user_id,
  role_id
)
VALUES (
  gen_random_uuid(),
  $1,
  $2
)
RETURNING *;
