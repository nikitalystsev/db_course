alter table if exists ss.retailer
    alter column id set not null,
    alter column title set not null,
    alter column address set not null,
    add unique (address),
    alter column phone_number set not null,
    add unique (phone_number),
    alter column fio_representative set not null,
    add unique (fio_representative);

alter table if exists ss.distributor
    alter column id set not null,
    alter column title set not null,
    alter column address set not null,
    add unique (address),
    alter column phone_number set not null,
    add unique (phone_number),
    alter column fio_representative set not null,
    add unique (fio_representative);

alter table if exists ss.manufacturer
    alter column id set not null,
    alter column title set not null,
    alter column address set not null,
    add unique (address),
    alter column phone_number set not null,
    add unique (phone_number),
    alter column fio_representative set not null,
    add unique (fio_representative);

alter table if exists ss.shop
    alter column id set not null,
    alter column retailer_id set not null,
    alter column title set not null,
    alter column address set not null,
    add unique (address),
    alter column phone_number set not null,
    add unique (phone_number),
    alter column fio_director set not null;

alter table if exists ss.product
    alter column id set not null,
    alter column retailer_id set not null,
    alter column distributor_id set not null,
    alter column manufacturer_id set not null,
    alter column name set not null,
    alter column categories set not null,
    alter column brand set not null,
    alter column compound set not null,
    alter column gross_mass set not null,
    alter column net_mass set not null,
    alter column package_type set not null,
    add check (gross_mass > 0 and net_mass > 0),
    add check (net_mass <= gross_mass);

alter table if exists ss.certificate_compliance
    alter column id set not null,
    alter column product_id set not null,
    alter column type set not null,
    alter column number set not null,
    alter column normative_document set not null,
    alter column status_compliance set not null,
    alter column registration_data set not null,
    alter column expiration_data set not null,
    add check (registration_data < expiration_data);

alter table ss.user
    alter column id set not null,
    alter column fio set not null,
    alter column phone_number set not null,
    add unique (phone_number),
    alter column password set not null,
    alter column registration_data set not null;

alter table ss.price
    alter column id set not null,
    alter column price set not null,
    alter column currency set not null,
    alter column setting_date set not null,
    add check (price > 0);

alter table ss.promotion
    alter column id set not null,
    alter column type set not null,
    alter column description set not null,
    alter column start_date set not null,
    alter column end_date set not null,
    add check (discount_size > 0 and discount_size < 100),
    add check (start_date < end_date);

alter table ss.sale_product
    alter column id set not null,
    alter column shop_id set not null,
    alter column product_id set not null,
    alter column price_id set not null,
    add check (rating >= 0 and rating <= 5),
    add unique (shop_id, product_id);

alter table ss.rating
    alter column id set not null,
    alter column user_id set not null,
    alter column sale_product_id set not null,
    alter column review set not null,
    alter column rating set not null,
    add check (rating >= 0 and rating <= 5);





