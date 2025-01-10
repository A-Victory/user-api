-- +goose Up
ALTER TABLE users
ADD COLUMN first_name VARCHAR(100),
ADD COLUMN last_name VARCHAR(100);

-- +goose Down
ALTER TABLE users
DROP COLUMN IF EXISTS first_name,
DROP COLUMN IF EXISTS last_name;
