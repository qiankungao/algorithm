
select DATE_FORMAT(order_date,'%Y-%m') as month, count(*) as order_count , count(distinct(customer_id)) as customer_count from Orders
where invoice>20
group by DATE_FORMAT(order_date,'%Y-%m')  ;