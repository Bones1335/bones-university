-- +goose Up
CREATE TABLE user_internships (
  user_internships_id UUID PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES users (users_id) ON DELETE CASCADE,
  internship_id UUID NOT NULL REFERENCES internships (internships_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE user_internships;
