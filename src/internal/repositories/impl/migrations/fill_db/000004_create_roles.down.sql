revoke administrator from admin_user;
revoke registered from registered_user;
revoke guest from guest_user;

drop user if exists admin_user;
drop user if exists registered_user;
drop user if exists guest_user;

revoke all privileges on all tables in schema ss from administrator;
revoke registered from administrator;
revoke guest from administrator;

revoke update on table ss.price from registered;
revoke insert on table ss.price from registered;
revoke insert on table ss.promotion from registered;
revoke insert on table ss.sale_product from registered;
revoke insert on table ss.rating from registered;
revoke insert on table ss.shop from registered;
revoke insert on table ss.product from registered;
revoke insert on table ss.price from registered;
revoke guest from registered;

revoke insert on table ss."user" from guest;
revoke select on all tables in schema ss from guest;

drop role if exists administrator;
drop role if exists registered;
drop role if exists guest;