-- name: CreateInternships :one
INSERT INTO internships (
  internships_id,
  business_name,
  num_spots,
  business_address,
  business_city,
  business_postal_code,
  business_state,
  business_country,
  business_phone_number,
  business_email,
  business_type
)
VALUES (
  gen_random_uuid(),
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8,
  $9,
  $10
)
RETURNING *;
