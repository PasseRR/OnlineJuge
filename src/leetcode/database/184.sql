/**
 * e表DepartmentId和d表Id非外键关系 可能e表存在DepartmentId而d表没有对应Id
 */
-- 方法一
SELECT
  d.NAME AS Department,
  e.NAME AS Employee,
  e.Salary
FROM
  Employee e
  LEFT JOIN Department d ON e.DepartmentId = d.Id
WHERE
  NOT EXISTS(
    SELECT 1
    FROM
      Employee
    WHERE
      DepartmentId = e.DepartmentId
      AND Salary > e.Salary
  )
  AND d.Id IS NOT NULL

-- 方法二
SELECT
  d.NAME AS Department,
  e.NAME AS Employee,
  t.Salary
FROM
  Employee e
  LEFT JOIN Department d ON e.DepartmentId = d.Id
  LEFT JOIN (
              SELECT
                DepartmentId,
                max(Salary) AS Salary
              FROM
                Employee
              GROUP BY
                DepartmentId
            ) AS t ON e.DepartmentId = t.DepartmentId AND e.Salary = t.Salary
WHERE
  t.DepartmentId IS NOT NULL
  AND d.NAME IS NOT NULL