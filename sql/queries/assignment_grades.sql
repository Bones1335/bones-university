-- name: CreateAssignmentGrades :one
INSERT INTO assignment_grades (
  assignment_grades_id,
  assignment_id,
  user_id
)
VALUES (
  gen_random_uuid(),
  $1,
  $2
)
RETURNING *;
