alter table if exists ss.retailer
    alter column id set not null,
    alter column id set default uuid_generate_v4(),
    alter column title set not null,
    alter column address set not null,
    add unique (address),
    alter column phone_number set not null,
    add unique (phone_number),
    alter column fio_representative set not null,
    add unique (fio_representative);

alter table if exists ss.distributor
    alter column id set not null,
    alter column id set default uuid_generate_v4(),
    alter column title set not null,
    alter column address set not null,
    add unique (address),
    alter column phone_number set not null,
    add unique (phone_number),
    alter column fio_representative set not null,
    add unique (fio_representative);

alter table if exists ss.manufacturer
    alter column id set not null,
    alter column id set default uuid_generate_v4(),
    alter column title set not null,
    alter column address set not null,
    add unique (address),
    alter column phone_number set not null,
    add unique (phone_number),
    alter column fio_representative set not null,
    add unique (fio_representative);

alter table if exists ss.shop
    alter column id set not null,
    alter column id set default uuid_generate_v4(),
    alter column retailer_id set not null,
    add foreign key (retailer_id) references retailer (id) on delete cascade on update cascade,
    alter column title set not null,
    alter column address set not null,
    add unique (address),
    alter column phone_number set not null,
    add unique (phone_number),
    alter column fio_director set not null;

alter table if exists ss.product
    alter column id set not null,
    alter column id set default uuid_generate_v4(),
    alter column retailer_id set not null,
    add foreign key (retailer_id) references retailer (id) on delete cascade on update cascade,
    alter column distributor_id set not null,
    add foreign key (distributor_id) references distributor (id) on delete cascade on update cascade,
    alter column manufacturer_id set not null,
    add foreign key (manufacturer_id) references manufacturer (id) on delete cascade on update cascade,
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
    alter column id set default uuid_generate_v4(),
    alter column product_id set not null,
    add foreign key (product_id) references product (id) on delete cascade on update cascade,
    alter column type set not null,
    alter column number set not null,
    alter column normative_document set not null,
    alter column status_compliance set not null,
    alter column registration_date set not null,
    alter column expiration_date set not null,
    add check (registration_date < expiration_date);

alter table ss.user
    alter column id set not null,
    alter column id set default uuid_generate_v4(),
    alter column fio set not null,
    alter column phone_number set not null,
    add unique (phone_number),
    alter column password set not null,
    alter column registration_date set not null;

alter table ss.promotion
    alter column id set not null,
    alter column id set default uuid_generate_v4(),
    alter column type set not null,
    alter column description set not null,
    alter column start_date set not null,
    alter column end_date set not null,
    add check (discount_size > 0 and discount_size < 100),
    add check (start_date < end_date);

alter table ss.sale_product
    alter column id set not null,
    alter column id set default uuid_generate_v4(),
    alter column shop_id set not null,
    add foreign key (shop_id) references shop (id) on delete cascade on update cascade,
    alter column product_id set not null,
    add foreign key (product_id) references product (id) on delete cascade on update cascade,
    add foreign key (promotion_id) references promotion (id) on delete cascade on update cascade,
    alter column price set not null,
    alter column currency set not null,
    alter column setting_date set not null,
    add check (avg_rating >= 0 and avg_rating <= 5),
    add check (price > 0),
    add unique (shop_id, product_id);

alter table ss.rating
    alter column id set not null,
    alter column id set default uuid_generate_v4(),
    alter column user_id set not null,
    add foreign key (user_id) references "user" (id) on delete cascade on update cascade,
    alter column sale_product_id set not null,
    add foreign key (sale_product_id) references sale_product (id) on delete cascade on update cascade,
    alter column review set not null,
    alter column rating set not null,
    add check (rating >= 0 and rating <= 5);

alter table ss.retailer_distributor
    alter column retailer_id set not null,
    add foreign key (retailer_id) references retailer (id) on delete cascade on update cascade,
    alter column distributor_id set not null,
    add foreign key (distributor_id) references distributor (id) on delete cascade on update cascade;

alter table ss.distributor_manufacturer
    alter column distributor_id set not null,
    add foreign key (distributor_id) references distributor (id) on delete cascade on update cascade,
    alter column manufacturer_id set not null,
    add foreign key (manufacturer_id) references manufacturer (id) on delete cascade on update cascade;



