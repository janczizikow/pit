SELECT name,class, season_id, COUNT(*)
FROM submissions AS s
GROUP BY name, class,season_id
HAVING COUNT(*)>1;
