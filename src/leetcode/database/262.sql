SELECT
  a.Request_at                                                      AS 'Day',
  CAST(FORMAT(IFNULL(b.canceled, 0) / a.total, 2) AS DECIMAL(3, 2)) AS 'Cancellation Rate'
FROM (
       SELECT
         Request_at,
         count(1) AS total
       FROM Trips t LEFT JOIN Users u ON t.Client_Id = u.Users_Id
       WHERE
         u.Banned = 'No'
       GROUP BY Request_at
     ) AS a -- 全部路程
  LEFT JOIN (
              SELECT
                Request_at,
                count(1) AS canceled
              FROM Trips t LEFT JOIN Users u ON t.Client_Id = u.Users_Id
              WHERE
                u.Banned = 'No' AND
                t.Status IN ('cancelled_by_client', 'cancelled_by_driver')
              GROUP BY Request_at
            ) AS b  -- 被取消的路程
    ON a.Request_at = b.Request_at
WHERE a.Request_at BETWEEN '2013-10-01' AND '2013-10-03'