CREATE TABLE users (
  id          SERIAL PRIMARY KEY,
  email       VARCHAR(64) NOT NULL,
  password    VARCHAR(64) NOT NULL,
  salt        VARCHAR(64) NOT NULL,
  created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at  TIMESTAMP DEFAULT NULL
);
