-- +goose Up
CREATE TABLE internships (
  internships_id UUID PRIMARY KEY,
  business_name TEXT NOT NULL,
  num_spots SMALLINT NOT NULL,
  business_address TEXT NOT NULL,
  business_city TEXT NOT NULL,
  business_postal_code INT NOT NULL,
  business_state TEXT NOT NULL,
  business_country TEXT NOT NULL,
  business_phone_number TEXT NOT NULL,
  business_email TEXT NOT NULL,
  business_type TEXT NOT NULL
);

-- +goose Down
DROP TABLE internships;
