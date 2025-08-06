-- name: CreateRole :one
INSERT INTO roles (
  roles_id,
  role_name
)
VALUES (
  gen_random_uuid(),
  $1
)
RETURNING *;

-- name: GetSingleRole :one 
SELECT * FROM roles
WHERE role_name = $1;
