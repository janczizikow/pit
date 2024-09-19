\copy submissions("name", class, tier, mode, build, video, duration, created_at, updated_at)
FROM 'data.csv'
WITH null as E'\'\'' CSV HEADER;
