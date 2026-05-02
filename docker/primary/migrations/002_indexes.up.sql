CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE INDEX idx_users_first_trgm ON users
USING gin (first_name gin_trgm_ops);

CREATE INDEX idx_users_second_trgm ON users
USING gin (second_name gin_trgm_ops);