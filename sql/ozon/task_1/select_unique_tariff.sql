SELECT pt.*
FROM pickpoint_tariff pt
         INNER JOIN (
    SELECT pickpoint_id, MAX(created_at) AS max_created_at
    FROM pickpoint_tariff
    GROUP BY pickpoint_id
) max_dates ON pt.pickpoint_id = max_dates.pickpoint_id AND pt.created_at = max_dates.max_created_at
ORDER BY created_at DESC ;