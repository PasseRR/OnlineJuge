CREATE FUNCTION getNthHighestSalary(N INT)
  RETURNS INT
  BEGIN
    DECLARE idx INT;
    SET idx = N - 1;
    RETURN (
      SELECT DISTINCT Salary
      FROM Employee
      ORDER BY Salary DESC
      LIMIT idx, 1
    );
  END