-- +goose Up 
CREATE TABLE years (
  years_id UUID PRIMARY KEY,
  year_in_school SMALLINT NOT NULL
);

-- +goose Down
DROP TABLE years;
