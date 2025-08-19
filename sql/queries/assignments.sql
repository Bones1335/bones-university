-- name: CreateAssignments :one
INSERT INTO assignments (
  assignments_id,
  created_at,
  updated_at,
  assignment_name,
  assignment_due_date,
  assignment_description,
  assignment_weight,
  course_id
)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING *;
