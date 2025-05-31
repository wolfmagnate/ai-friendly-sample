CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  created_at BIGINT NOT NULL,
  updated_at BIGINT NOT NULL,
  user_uid VARCHAR(255) NOT NULL
);
