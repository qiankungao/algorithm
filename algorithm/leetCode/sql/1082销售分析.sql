select seller_id  from `Sales` group by seller_id having sum(price)=(

select sum(price) as sum  from `Sales` group by seller_id order by sum desc limit 1

);
