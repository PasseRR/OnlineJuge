/**
 * 学生和课程存在重复
 */
SELECT class
FROM courses
GROUP BY class
HAVING COUNT(DISTINCT student) >= 5;