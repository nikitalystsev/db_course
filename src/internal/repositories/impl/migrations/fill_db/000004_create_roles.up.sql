create role guest;

grant usage on schema ss to guest;
grant select on all tables in schema ss to guest;
grant insert on table ss.user to guest;


create role registered;

grant guest to registered;

grant insert on table ss.promotion to registered;
grant insert on table ss.sale_product to registered;
grant insert on table ss.rating to registered;
grant insert on table ss.shop to registered;
grant insert on table ss.product to registered;
grant insert on table ss.retailer to registered;
grant insert on table ss.distributor to registered;
grant insert on table ss.manufacturer to registered;
grant insert on table ss.retailer_distributor to registered;
grant insert on table ss.distributor_manufacturer to registered;

grant update on table ss.sale_product to registered;


create role administrator;

grant registered to administrator;

grant all privileges on all tables in schema ss to administrator;


create user guest_user with password 'guest';

create user registered_user with password 'registered';

create user admin_user with password 'admin';


grant guest to guest_user;

grant registered to registered_user;

grant administrator to admin_user;
