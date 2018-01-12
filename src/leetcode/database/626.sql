-- 方法一
SELECT
  s1.id,
  IFNULL(
    (
      SELECT s2.student
      FROM
        seat s2
      WHERE
        s2.id = (
          CASE
          WHEN s1.id % 2 = 1
            THEN s1.id + 1
          ELSE s1.id - 1
          END
        )
    ),
    s1.student
  ) AS student
FROM
  seat s1
ORDER BY
  s1.id;

-- 方法二
SELECT
  s1.id,
  COALESCE(s2.student, s1.student) AS student
FROM
  seat s1 LEFT JOIN seat s2 ON s2.id = ((s1.id + 1) ^ 1) - 1
ORDER BY s1.id;