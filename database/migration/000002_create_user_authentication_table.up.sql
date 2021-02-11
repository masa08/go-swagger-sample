CREATE TABLE IF NOT EXISTS user_authentication
(
  user_id    serial PRIMARY KEY NOT NULL REFERENCES users (id),
  login_id   text NOT NULL UNIQUE,
  password   text NOT NULL UNIQUE,
  created_at timestamp(0) without time zone NOT NULL default now(),
  updated_at timestamp(0) without time zone NOT NULL default now()
);
