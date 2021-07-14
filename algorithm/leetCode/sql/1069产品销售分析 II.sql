select p.product_id,sum(s.quantity) as total_quantity from Product as p

left join Sales as s on p.product_id = s.product_id

where s.product_id is not null

group by p.product_id ;