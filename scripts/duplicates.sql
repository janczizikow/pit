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

--        name       |    class    | tier | duration | season_id |   mode   
-- ------------------+-------------+------+----------+-----------+----------
--  jacky0835        | sorcerer    |  140 |      777 |         5 | softcore
--  Storm            | sorcerer    |  141 |      784 |         5 | softcore
--  DraGuNoVxTT      | sorcerer    |  134 |      857 |         5 | softcore
--  배반             | barbarian   |  135 |      857 |         4 | softcore
--  光头爸爸official | barbarian   |  140 |      857 |         5 | softcore
--  Forestbog        | rogue       |  146 |      897 |         5 | softcore
--  考考你丫         | rogue       |  134 |      897 |         4 | softcore
--  Skydro           | sorcerer    |  156 |      897 |         5 | softcore
--  ChrisGrandPa     | sorcerer    |  153 |      857 |         5 | softcore
--  Skydro           | sorcerer    |  156 |      897 |         5 | softcore
--  Storm            | sorcerer    |  141 |      784 |         5 | softcore
--  Winzik           | necromancer |  150 |      897 |         5 | softcore
--  FP               | rogue       |  146 |      857 |         5 | softcore
--  RR1              | sorcerer    |  151 |      777 |         5 | softcore
--  秋葉-Akiha-      | necromancer |  145 |      857 |         5 | softcore
--  冲锋的米卡       | necromancer |  130 |      857 |         5 | softcore
--  FARFAR           | druid       |  145 |      777 |         5 | softcore
--  너즐             | necromancer |  127 |      777 |         4 | softcore
--  Mydas            | rogue       |  123 |      857 |         4 | softcore
--  Mental           | rogue       |  119 |      857 |         4 | softcore
--  Pxx              | barbarian   |  160 |      897 |         4 | softcore
--  Sonya            | barbarian   |  135 |      857 |         4 | hardcore
--  RR1              | sorcerer    |  151 |      777 |         5 | softcore
--  ChrisGrandPa     | sorcerer    |  153 |      857 |         5 | softcore
-- (24 rows)