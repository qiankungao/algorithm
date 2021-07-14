select p.product_name,sum(o.unit) as unit from Orders as o left join Products as p on p.product_id = o.product_id
where  year(o.order_date)=2020 and month(o.order_date) = 2
group by o.product_id
having sum(o.unit)>=100 ;