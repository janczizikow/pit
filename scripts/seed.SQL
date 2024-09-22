\COPY seasons(id, "name", pit, "start", "end")
FROM 'seeds/seasons.csv'
CSV HEADER DELIMITER ',';

\COPY submissions("name", class, tier, mode, build, video, duration, verified, created_at, updated_at)
FROM 'data.csv'
WITH null as E'\'\'' CSV HEADER DELIMITER ',';
