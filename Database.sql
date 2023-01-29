-- Number 1
-- The test case is not clear
SELECT x.A, X.B, y.E 
FROM y
INNER JOIN x ON x.C = y.D


-- Number 2
-- Hi kak Dicky, the given data type for tanggal is not defined so I'll make 2 query for 2 different data type

-- If it's an date
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