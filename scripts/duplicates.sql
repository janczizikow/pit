SELECT name,class,tier,duration,season_id,mode
FROM submissions
WHERE duration IN (
        SELECT duration FROM (
                SELECT name,tier, duration,season_id, COUNT(*)
                FROM submissions
                GROUP BY name,tier, duration, season_id
                HAVING COUNT(*)>1

        ) AS a
);
