CREATE TABLE IF NOT EXISTS submissions (
  id SERIAL PRIMARY KEY,
  "name" varchar(255) NOT NULL DEFAULT '',
  class varchar(255) NOT NULL DEFAULT '',
  tier int NOT NULL DEFAULT 1 CHECK (tier > 0),
  build varchar(255) NOT NULL DEFAULT '',
  video varchar(255) NOT NULL DEFAULT '',
  duration int NOT NULL DEFAULT 0,
  created_at timestamp(0) with time zone NOT NULL DEFAULT now(),
  updated_at timestamp(0) with time zone NOT NULL DEFAULT now()
);
