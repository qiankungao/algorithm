select d.id,
sum(if(d.month='Jan',revenue,null)) as Jan_Revenue ,

sum(if(d.month='Feb',revenue,null)) as Feb_Revenue,

sum(if(d.month='Mar',revenue,null)) as Mar_Revenue,
sum(if(d.month='Apr',revenue,null)) as Apr_Revenue,
sum(if(d.month='May',revenue,null)) as May_Revenue,
sum(if(d.month='Jun',revenue,null)) as Jun_Revenue,
sum(if(d.month='Jul',revenue,null)) as Jul_Revenue,
sum(if(d.month='Aug',revenue,null)) as Aug_Revenue,
sum(if(d.month='Sep',revenue,null)) as Sep_Revenue,
sum(if(d.month='Oct',revenue,null)) as Oct_Revenue,
sum(if(d.month='Nov',revenue,null)) as Nov_Revenue,
sum(if(d.month='Dec',revenue,null)) as Dec_Revenue

from `Department` as d group by d.id;



select d.id,
sum(case d.month when 'Jan' then revenue ELSE null END) as Jan_Revenue ,

sum(case d.month  when 'Feb' then revenue else null end) as Feb_Revenue,

sum(case d.month  when 'Mar' then revenue else null end) as Mar_Revenue,
sum(case d.month  when 'Apr' then revenue else null end)  as Apr_Revenue,
sum(case d.month  when 'May' then revenue else null end)  as May_Revenue,
sum(case d.month  when 'Jun' then revenue else null end)  as Jun_Revenue,
sum(case d.month  when 'Jul' then revenue else null end)  as Jul_Revenue,
sum(case d.month  when 'Aug' then revenue else null end)  as Aug_Revenue,
sum(case d.month  when 'Sep' then revenue else null end) as Sep_Revenue,
sum(case d.month  when 'Oct' then revenue else null end) as Oct_Revenue,
sum(case d.month  when 'Nov' then revenue else null end)  as Nov_Revenue,
sum(case d.month  when 'Dec' then revenue else null end) as Dec_Revenue
from `Department` as d group by d.id;