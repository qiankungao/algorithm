select e.employee_id,e2.team_size from Employee as e

join (select team_id,count(*) as team_size from `Employee` group by team_id) as e2

on e.`team_id` = e2.`team_id` ;



select e1.employee_id, count(*) team_size
from employee e1 left join employee e2
on e1.team_id = e2.team_id
group by e1.employee_id;