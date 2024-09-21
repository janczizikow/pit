\COPY submissions("name", class, tier, mode, build, video, duration, verified, created_at, updated_at)
FROM 'data.csv'
WITH null as E'\'\'' CSV HEADER DELIMITER ',';
