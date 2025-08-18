-- +goose Up
CREATE TABLE degrees_courses (
  degrees_courses_id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  degree_id UUID NOT NULL REFERENCES degree_programs (degrees_id) ON DELETE CASCADE,
  course_id UUID NOT NULL REFERENCES courses (courses_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE degrees_courses;
