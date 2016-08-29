CREATE FUNCTION getNthHighestSalary(N INT)
  RETURNS INT
  BEGIN
    -- 分页索引位置 mysql从0开始
    DECLARE idx INT;
    SET idx = N - 1;
    RETURN (
      SELECT DISTINCT Salary
      FROM Employee
      ORDER BY Salary DESC
      LIMIT idx, 1
    );
  END