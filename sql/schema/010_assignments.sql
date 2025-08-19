-- +goose Up
CREATE TABLE assignments (
  assignments_id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  assignment_name TEXT NOT NULL,
  assignment_due_date TIMESTAMP NOT NULL,
  assignment_description TEXT NOT NULL,
  assignment_weight INT NOT NULL,
  course_id UUID NOT NULL REFERENCES courses (courses_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE assignments;
