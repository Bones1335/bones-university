-- +goose Up
CREATE TABLE course_enrollment (
  enrollment_id UUID PRIMARY KEY,
  final_grade NUMERIC(5,2) NOT NULL DEFAULT 0.00,
  course_id UUID NOT NULL REFERENCES courses (courses_id) ON DELETE CASCADE,
  user_id UUID NOT NULL REFERENCES users (users_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE course_enrollment;
