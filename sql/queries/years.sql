-- name: CreateYears :one
INSERT INTO years (
  years_id,
  year_in_school
)
VALUES (
  gen_random_uuid(),
  $1
)
RETURNING *;


