SELECT w1.Id
FROM
  Weather w1, Weather w2
WHERE w1.Date = DATE_ADD(w2.Date, INTERVAL 1 DAY)
      AND w1.Temperature > w2.Temperature