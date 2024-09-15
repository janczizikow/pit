COPY submissions("name", class, tier, mode, build, video, duration, created_at, updated_at)
FROM $(pwd)
DELIMITER ','
CSV HEADER;
