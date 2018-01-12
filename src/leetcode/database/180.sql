-- 方法一
SELECT DISTINCT l1.Num AS ConsecutiveNums
FROM
  Logs l1,
  Logs l2,
  Logs l3
WHERE
  l1.Id = l2.Id - 1
  AND l2.Id = l3.Id - 1
  AND l1.Num = l2.Num
  AND l2.Num = l3.Num;

-- 方法二
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
WHERE A.CNT >= 3;
