-- name: CreateCourses :one 
INSERT INTO courses (
  courses_id,
  created_at,
  updated_at,
  course_code,
  course_name,
  course_description,
  course_professor_id
)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1,
  $2,
  $3,
  $4
)
RETURNING *;
