begin;

select plan(2);

insert into ss.rating (user_id, sale_product_id, review, rating)

values ('362b79f6-d671-404a-b1a0-5a655aebc1b6', '53b7ff0a-b5cd-4107-85af-1617b137aa23', 'great', 5);

select is(
               (select avg_rating from ss.sale_product where id = '53b7ff0a-b5cd-4107-85af-1617b137aa23'),
               5.0
       );

delete from ss.rating
where user_id = '362b79f6-d671-404a-b1a0-5a655aebc1b6' and sale_product_id = '53b7ff0a-b5cd-4107-85af-1617b137aa23';

select is(
               (select avg_rating from ss.sale_product where id = '53b7ff0a-b5cd-4107-85af-1617b137aa23'),
               null
       );

select *
from finish();

rollback;