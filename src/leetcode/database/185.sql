/**
 * Hint
 * 1. e表DepartmentId和d表Id非外键关系 可能e表存在DepartmentId而d表没有对应Id
 * 2. 同一个部门同样工资按照相同排名计算
 * 3. 部门排序、工资降序
 */
-- 方法一
SELECT
  d.Name AS Department,
  e.Name AS Employee,
  e.Salary
FROM
  Employee e LEFT JOIN Department d ON e.DepartmentId = d.Id
WHERE
  d.Id IS NOT NULL
  AND (
        SELECT COUNT(DISTINCT (Salary))
        FROM
          Employee
        WHERE
          DepartmentId = e.DepartmentId
          AND Salary >= e.Salary
      ) < 4
ORDER BY
  d.Id ASC, e.Salary DESC;

-- 方法二
SELECT
  d.Name AS Department,
  e.Name AS Employee,
  e.Salary
FROM (
       SELECT
         e1.Name,
         e1.Salary,
         e1.DepartmentId
       FROM
         Employee e1 JOIN Employee e2
           ON e1.DepartmentId = e2.DepartmentId AND e1.Salary <= e2.Salary
       GROUP BY e1.Id
       HAVING COUNT(DISTINCT (e2.Salary)) < 4
     ) AS e LEFT JOIN Department d ON e.DepartmentId = d.Id
WHERE d.Id IS NOT NULL
ORDER BY d.Id, e.Salary DESC;