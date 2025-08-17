-- +goose Up
CREATE TABLE courses (
  courses_id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  course_code TEXT NOT NULL,
  course_name TEXT NOT NULL,
  course_description TEXT NOT NULL,
  course_professor_id UUID NOT NULL REFERENCES users (users_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE courses;
