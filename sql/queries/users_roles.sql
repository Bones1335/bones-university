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

-- name: GetUsersRole :one
SELECT roles.role_name FROM roles
INNER JOIN users_roles ON users_roles.role_id = roles.roles_id
WHERE users_roles.user_id = $1;

-- name: UpdateUsersRole :one
UPDATE users_roles SET role_id = $1
WHERE user_id = $2
RETURNING *;
