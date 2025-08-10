-- name: CreateUser :one
INSERT INTO users (
    users_id,
    created_at,
    updated_at,
    last_name,
    first_name,
    username,
    personal_email,
    university_email,
    password
)
Values (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: Login :one
SELECT * FROM users
WHERE username = $1;

-- name: UpdateUser :one
UPDATE users SET last_name = $2, first_name = $3, personal_email = $4, password = $5
WHERE users_id = $1
RETURNING *;
