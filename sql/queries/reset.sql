-- name: ResetUsers :exec
DELETE FROM users;

-- name: ResetRoles :exec
DELETE FROM roles;
