-- +goose Up 
CREATE TABLE assignment_grades (
  assignment_grades_id UUID PRIMARY KEY,
  assingmnet_grade NUMERIC(5,2) NOT NULL DEFAULT 0.00,
  user_id UUID NOT NULL REFERENCES users (users_id) ON DELETE CASCADE,
  assignment_id UUID NOT NULL REFERENCES assignments (assignments_id) ON DELETE CASCADE
);

-- +goose Down 
DROP TABLE assignment_grades;
