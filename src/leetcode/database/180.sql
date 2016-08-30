SELECT DISTINCT Num AS ConsecutiveNums
FROM (
       SELECT
         Num,
         CASE
         WHEN @prev = Num
           THEN @count := @count + 1
         WHEN (@prev := Num) IS NOT NULL
           THEN @count := 1
         END AS CNT
       FROM Logs, (SELECT @prev := NULL) X
       ORDER BY Id
     ) AS A
WHERE A.CNT >= 3
