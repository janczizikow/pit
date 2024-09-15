CREATE TYPE class AS ENUM ('barbarian', 'druid', 'necromancer', 'rogue', 'sorcerer');
CREATE TYPE mode AS ENUM ('softcore', 'hardcore');
CREATE TABLE IF NOT EXISTS submissions (
  id SERIAL PRIMARY KEY,
  "name" varchar(255) NOT NULL DEFAULT '',
  "class" class NOT NULL,
  tier int NOT NULL DEFAULT 1 CHECK (tier > 0),
  "mode" mode NOT NULL DEFAULT 'softcore',
  build text NOT NULL DEFAULT '',
  video text NOT NULL DEFAULT '',
  duration int NOT NULL DEFAULT 0,
  created_at timestamp(0) with time zone NOT NULL DEFAULT now(),
  updated_at timestamp(0) with time zone NOT NULL DEFAULT now()
);
