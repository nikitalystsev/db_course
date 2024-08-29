create or replace function ss.get_sale_product_avg_rating(_sale_product_id uuid)
    returns numeric
as
$$
declare
    cnt_ratings int;
    avg_rating  numeric;

begin
    select count(*)
    into cnt_ratings
    from ss.rating as r
    where r.sale_product_id = _sale_product_id;

    if cnt_ratings != 0 then
        select avg(r.rating)
        into avg_rating
        from ss.rating as r
        where r.sale_product_id = _sale_product_id;

        return avg_rating;
    else
        return null;
    end if;
end;
$$
    language plpgsql;

create or replace function ss.update_sale_product_avg_rating()
    returns trigger
as
$$
begin
    if tg_op = 'INSERT' then

        update ss.sale_product
        set avg_rating = ss.get_sale_product_avg_rating(new.sale_product_id)
        where id = new.sale_product_id;

        return new;
    elseif tg_op = 'DELETE' then

        update ss.sale_product
        set avg_rating = ss.get_sale_product_avg_rating(old.sale_product_id)
        where id = old.sale_product_id;

        return old;
    end if;
end;
$$
    language plpgsql;

create or replace trigger insert_sale_product_avg_rating
    after insert
    on ss.rating
    for each row
execute function ss.update_sale_product_avg_rating();

create or replace trigger delete_sale_product_avg_rating
    after delete
    on ss.rating
    for each row
execute function ss.update_sale_product_avg_rating();