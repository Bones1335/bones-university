-- name: CreateDegreePrograms :one
INSERT INTO degree_programs (
  degrees_id, 
  created_at,
  updated_at,
  degree_name,
  degree_level,
  degree_department,
  degree_duration
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

-- name: GetDegrees :many 
SELECT * FROM degree_programs;
