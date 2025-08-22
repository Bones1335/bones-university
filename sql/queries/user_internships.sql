-- name: CreateUserInternships :one
INSERT INTO user_internships (
  user_internships_id,
  user_id,
  internship_id
)
VALUES (
  gen_random_uuid(),
  $1,
  $2
)
RETURNING *;
