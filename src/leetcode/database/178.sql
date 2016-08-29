SELECT
  Score,
  ((SELECT COUNT(DISTINCT Score)
    FROM Scores
    WHERE Score > a.Score) + 1) AS Rank
FROM Scores a
ORDER BY Score DESC