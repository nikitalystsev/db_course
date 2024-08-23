create or replace function ss.update_avg_rating()
    returns trigger
as
$$
declare
    new_avg_rating numeric;

begin
    select avg(rating)
    into new_avg_rating
    from ss.rating
    where sale_product_id = new.sale_product_id;

    update ss.sale_product set avg_rating = new_avg_rating where id = new.sale_product_id;

    return null;
end;
$$
    language plpgsql;


create or replace trigger check_avg_rating
    after insert
    on ss.rating
    for each row
execute function ss.update_avg_rating();