\COPY seasons(id, "name", pit, "start", "end")
FROM 'seeds/seasons.csv'
CSV HEADER DELIMITER ',';

\COPY submissions("name", class, tier, mode, video, build, duration, verified, season_id)
FROM 'seeds/season4.csv'
WITH null as E'\'\'' CSV HEADER DELIMITER ',';

\COPY submissions("name",class,tier,mode,video,build,duration,verified,season_id)
FROM 'seeds/helltides.csv'
WITH null as E'\'\'' CSV HEADER DELIMITER ',';

\COPY submissions("name", class, tier, mode, build, video, duration, verified, season_id, created_at, updated_at)
FROM 'data.csv'
WITH null as E'\'\'' CSV HEADER DELIMITER ',';
