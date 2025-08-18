-- name: CreateDegreesCourse :one
INSERT INTO degrees_courses (
  degrees_courses_id,
  created_at,
  updated_at,
  degree_id,
  course_id
)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1,
  $2
)
RETURNING *;
