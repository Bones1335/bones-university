-- name: CreateStudentsProgram :one
INSERT INTO users_programs (
  users_program_id,
  cohort_year,
  start_date,
  student_id,
  program_id,
  academic_year_id
)
VALUES (
  gen_random_uuid(),
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING *;
