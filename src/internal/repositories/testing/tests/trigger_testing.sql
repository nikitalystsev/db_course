begin;

select plan(1);

insert into ss.rating (user_id, sale_product_id, review, rating)

values ('362b79f6-d671-404a-b1a0-5a655aebc1b6', '53b7ff0a-b5cd-4107-85af-1617b137aa23', 'класс', 5);

select is(
               ss.get_sale_product_avg_rating('53b7ff0a-b5cd-4107-85af-1617b137aa23'),
               5.0
       );

select *
from finish();

rollback;