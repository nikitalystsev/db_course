begin;

select plan(5);

select is(
               ss.get_sale_product_avg_rating(
                       '53b7ff0a-b5cd-4107-85af-1617b137aa23'
               ),
               null
       );

insert into ss.rating (user_id, sale_product_id, review, rating)
values ('362b79f6-d671-404a-b1a0-5a655aebc1b6', '53b7ff0a-b5cd-4107-85af-1617b137aa23', 'great', 5);

insert into ss.rating (user_id, sale_product_id, review, rating)
values ('8d9b001f-5760-4c40-bc60-988e0ca54d18', '53b7ff0a-b5cd-4107-85af-1617b137aa23', 'bad', 2);


select is(
               ss.get_sale_product_avg_rating(
                       '53b7ff0a-b5cd-4107-85af-1617b137aa23'
               ),
               3.5
       );

select throws_ok(
               'ss.get_sale_product_avg_rating('')'
       );

select throws_ok(
               'ss.get_sale_product_avg_rating()'
       );

select throws_ok(
               'ss.get_sale_product_avg_rating(''tsgsddgd'')'
       );

select *
from finish();

rollback;