create extension if not exists "uuid-ossp";

create table if not exists ss.retailer
(
    id                 uuid primary key default uuid_generate_v4(),
    title              text,
    address            text,
    phone_number       varchar(30),
    fio_representative text
);

create table if not exists ss.distributor
(
    id                 uuid primary key default uuid_generate_v4(),
    title              text,
    address            text,
    phone_number       varchar(30),
    fio_representative text
);

create table if not exists ss.manufacturer
(
    id                 uuid primary key default uuid_generate_v4(),
    title              text,
    address            text,
    phone_number       varchar(30),
    fio_representative text
);


create table if not exists ss.shop
(
    id           uuid primary key default uuid_generate_v4(),
    retailer_id  uuid,
    title        text,
    address      text,
    phone_number varchar(30),
    fio_director text,
    foreign key (retailer_id) references retailer (id) on delete cascade on update cascade
);

create table if not exists ss.product
(
    id              uuid primary key default uuid_generate_v4(),
    retailer_id     uuid,
    distributor_id  uuid,
    manufacturer_id uuid,
    name            text,
    categories      text,
    brand           text,
    compound        text,
    gross_mass      numeric,
    net_mass        numeric,
    package_type    text,
    foreign key (retailer_id) references retailer (id) on delete cascade on update cascade,
    foreign key (distributor_id) references distributor (id) on delete cascade on update cascade,
    foreign key (manufacturer_id) references manufacturer (id) on delete cascade on update cascade
);

create table if not exists ss.certificate_compliance
(
    id                 uuid primary key default uuid_generate_v4(),
    product_id         uuid,
    type               text,
    number             text,
    normative_document text,
    status_compliance  boolean,
    registration_data  date,
    expiration_data    date,
    foreign key (product_id) references product (id) on delete cascade on update cascade
);

create table if not exists ss.user
(
    id                uuid primary key default uuid_generate_v4(),
    fio               text,
    phone_number      varchar(30),
    password          text,
    registration_data date
);

create table if not exists ss.price
(
    id           uuid primary key default uuid_generate_v4(),
    price        numeric,
    currency     text,
    setting_date date
);

create table if not exists ss.promotion
(
    id            uuid primary key default uuid_generate_v4(),
    type          text,
    description   text,
    discount_size numeric,
    start_date    date,
    end_date      date
);

create table if not exists ss.sale_product
(
    id           uuid primary key default uuid_generate_v4(),
    shop_id      uuid,
    product_id   uuid,
    price_id     uuid,
    promotion_id uuid,
    rating       float,
    foreign key (shop_id) references shop (id) on delete cascade on update cascade,
    foreign key (product_id) references product (id) on delete cascade on update cascade,
    foreign key (price_id) references price (id) on delete cascade on update cascade,
    foreign key (promotion_id) references promotion (id) on delete cascade on update cascade
);

create table if not exists ss.rating
(
    id              uuid primary key default uuid_generate_v4(),
    user_id         uuid,
    sale_product_id uuid,
    review          text,
    rating          numeric,
    foreign key (user_id) references "user" (id) on delete cascade on update cascade,
    foreign key (sale_product_id) references sale_product (id) on delete cascade on update cascade
);

create table if not exists ss.retailer_distributor
(
    retailer_id    uuid,
    distributor_id uuid,
    primary key (retailer_id, distributor_id),
    foreign key (retailer_id) references retailer (id) on delete cascade on update cascade,
    foreign key (distributor_id) references distributor (id) on delete cascade on update cascade
);

create table if not exists ss.distributor_manufacturer
(
    distributor_id  uuid,
    manufacturer_id uuid,
    primary key (distributor_id, manufacturer_id),
    foreign key (distributor_id) references distributor (id) on delete cascade on update cascade,
    foreign key (manufacturer_id) references manufacturer (id) on delete cascade on update cascade
);