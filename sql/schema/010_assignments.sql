-- +goose Up
CREATE TABLE assignments (
  assignments_id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  assignment_name TEXT NOT NULL,
  assignment_due_date TIMESTAMP NOT NULL,
  assignment_description TEXT NOT NULL
);

-- +goose Down
DROP TABLE assignments;
