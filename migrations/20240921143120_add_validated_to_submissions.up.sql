ALTER TABLE submissions ADD COLUMN IF NOT EXISTS verified boolean NOT NULL DEFAULT false;
