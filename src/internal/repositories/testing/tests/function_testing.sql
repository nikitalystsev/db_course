begin;

select plan(1);

select is(
               ss.get_sale_product_avg_rating('53b7ff0a-b5cd-4107-85af-1617b137aa23'),
               null
       );

select *
from finish();

rollback;