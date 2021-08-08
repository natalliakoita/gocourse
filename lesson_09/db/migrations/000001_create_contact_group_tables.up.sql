BEGIN;

CREATE TABLE IF NOT EXISTS groups (
  id serial PRIMARY KEY,
  name text NOT NULL
);

CREATE TABLE IF NOT EXISTS contacts (
  id serial PRIMARY KEY,
  name text NOT NULL,
  number text NOT NULL UNIQUE,
  group_id integer,
  FOREIGN KEY ("group_id") REFERENCES "groups"("id") ON DELETE CASCADE
);

INSERT INTO groups
  ( name )
VALUES
  ('News'), 
  ('Sport'), 
  ('Music'),
  ('Cook');

COMMIT;