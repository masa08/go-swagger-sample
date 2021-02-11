CREATE TABLE IF NOT EXISTS users
(
  id          serial PRIMARY KEY,
  email       text NOT NULL UNIQUE,
  first_name  text NOT NULL,
  last_name   text NOT NULL,
  phone       text NOT NULL UNIQUE,
  user_status integer NOT NULL ,
  user_name   text NOT NULL,
  created_at  timestamp(0) without time zone NOT NULL default now(),
  updated_at  timestamp(0) without time zone NOT NULL default now(),
  deleted_at  timestamp(0) without time zone default NULL,
);