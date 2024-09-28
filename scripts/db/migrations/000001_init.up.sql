CREATE TABLE IF NOT EXISTS platform_registry (
  id varchar(36) NOT NULL UNIQUE PRIMARY KEY,
  project_name varchar(255) NOT NULL,
  repository varchar(255) NOT NULL,
  username varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  "created_at" varchar NOT NULL,
  "deleted_at" varchar NOT NULL
);