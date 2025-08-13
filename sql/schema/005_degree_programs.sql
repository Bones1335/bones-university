-- +goose Up
CREATE TABLE degree_programs (
  degrees_id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  degree_name TEXT UNIQUE NOT NULL,
  degree_level TEXT NOT NULL,
  degree_department TEXT NOT NULL,
  degree_duration SMALLINT NOT NULL
);

-- +goose Down
DROP TABLE degree_programs;
