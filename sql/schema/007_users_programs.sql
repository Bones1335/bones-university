-- +goose Up
CREATE TABLE users_programs (
  users_program_id UUID PRIMARY KEY,
  cohort_year INT NOT NULL,
  start_date TIMESTAMP NOT NULL,
  student_id UUID NOT NULL REFERENCES users (users_id) ON DELETE CASCADE,
  program_id UUID NOT NULL REFERENCES degree_programs (degrees_id) ON DELETE CASCADE,
  academic_year_id UUID NOT NULL REFERENCES years (years_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE users_programs;
