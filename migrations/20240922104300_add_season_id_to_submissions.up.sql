ALTER TABLE submissions ADD COLUMN IF NOT EXISTS season_id int;
ALTER TABLE submissions ADD CONSTRAINT fk_season FOREIGN KEY (season_id) REFERENCES seasons(id);
