drop trigger if exists delete_sale_product_avg_rating on ss.rating;
drop trigger if exists insert_sale_product_avg_rating on ss.rating;

drop function if exists ss.update_sale_product_avg_rating;
drop function if exists ss.get_sale_product_avg_rating;