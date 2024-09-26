ALTER TYPE class RENAME TO class_old;
CREATE TYPE class AS ENUM ('barbarian', 'druid', 'necromancer', 'rogue', 'sorcerer');
ALTER TABLE submissions ALTER COLUMN class TYPE class USING class::text::class;
DROP TYPE class_old;
