-- name: CreateCourseEnrollment :one
INSERT INTO course_enrollment (
  enrollment_id,
  course_id,
  user_id
)
VALUES (
  gen_random_uuid(),
  $1,
  $2
)
RETURNING *;
