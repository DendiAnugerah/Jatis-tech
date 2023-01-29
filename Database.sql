-- Number 1
-- The test case is not clear
SELECT y.D, y.E 
FROM y
INNER JOIN x ON x.C = y.D


-- Number 2
-- Because the data type of tanggal is not defined in the problem, I'll make 2 query for 2 different data type

-- It's if the data type of tanggal is date kak dicky
SELECT YEAR(Tanggal) as Tahun,
       SUM(CASE WHEN Choice = 'A' THEN 1 ELSE 0 END) as A,
       SUM(CASE WHEN Choice = 'B' THEN 1 ELSE 0 END) as B,
       SUM(CASE WHEN Choice = 'C' THEN 1 ELSE 0 END) as C
FROM jatis
GROUP BY YEAR(Tanggal);

-- If it's an varchar, here the example
SELECT SUBSTRING_INDEX(Tanggal, '/', 3) as Tahun,
       SUM(CASE WHEN Choice = 'A' THEN 1 ELSE 0 END) as A,
       SUM(CASE WHEN Choice = 'B' THEN 1 ELSE 0 END) as B,
       SUM(CASE WHEN Choice = 'C' THEN 1 ELSE 0 END) as C
FROM jatis
GROUP BY SUBSTRING_INDEX(Tanggal, '/', 3);