select c.customer_id,c.name from Customers as c where c.customer_id in( select o.customer_id  from Orders as o
left join Product as p on o.product_id = p.product_id
where o.order_date between '2020-06-01' and '2020-07-31'
group by o.customer_id
having sum(case when o.order_date between '2020-06-01' and '2020-06-30' then o.quantity*p.price else 0 end
        )>=100
        and
        sum(case when o.order_date between '2020-07-01' and '2020-07-31' then o.quantity*p.price else 0 end
        )>=100
);