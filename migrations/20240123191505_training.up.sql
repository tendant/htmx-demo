CREATE TABLE IF NOT EXISTS training (
  uuid uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
  deleted_at timestamp
);
