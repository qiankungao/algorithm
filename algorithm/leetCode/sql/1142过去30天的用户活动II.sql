select ifnull(round( avg(count),2),0) as average_sessions_per_user from(select count(distinct(session_id)) as count from Activity where datediff('2019-7-27',activity_date)<30 group by user_id) as a;
SELECT IFNULL(ROUND(COUNT(DISTINCT session_id) / COUNT(DISTINCT user_id), 2), 0) AS average_sessions_per_user
FROM Activity
WHERE DATEDIFF('2019-07-27', activity_date) < 30

