SELECT e2.Name AS Employee
FROM Employee e1 LEFT JOIN Employee e2
    ON e1.Id = e2.ManagerId AND e2.ManagerId IS NOT NULL
WHERE e1.Salary < e2.Salary