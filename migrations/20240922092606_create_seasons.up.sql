CREATE TABLE IF NOT EXISTS seasons (
  id SERIAL PRIMARY KEY,
  name text NOT NULL DEFAULT '' UNIQUE,
  "start" timestamp(0) with time zone NOT NULL,
  "end" timestamp(0) with time zone,
  pit boolean DEFAULT false,
  created_at timestamp(0) with time zone NOT NULL DEFAULT now(),
  updated_at timestamp(0) with time zone NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS seasons_pit_idx ON seasons (pit);
