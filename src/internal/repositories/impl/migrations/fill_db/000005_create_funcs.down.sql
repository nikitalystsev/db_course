drop trigger if exists recount_sale_product_avg_rating on ss.rating;

drop function if exists ss.recount_sale_product_avg_rating;
drop function if exists ss.get_sale_product_avg_rating;