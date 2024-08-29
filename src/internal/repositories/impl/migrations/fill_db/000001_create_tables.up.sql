create extension if not exists "uuid-ossp";

create table if not exists ss.retailer
(
    id                 uuid primary key,
    title              text,
    address            text,
    phone_number       varchar(30),
    fio_representative text
);

create table if not exists ss.distributor
(
    id                 uuid primary key,
    title              text,
    address            text,
    phone_number       varchar(30),
    fio_representative text
);

create table if not exists ss.manufacturer
(
    id                 uuid primary key,
    title              text,
    address            text,
    phone_number       varchar(30),
    fio_representative text
);


create table if not exists ss.shop
(
    id           uuid primary key,
    retailer_id  uuid,
    title        text,
    address      text,
    phone_number varchar(30),
    fio_director text
);

create table if not exists ss.product
(
    id              uuid primary key,
    retailer_id     uuid,
    distributor_id  uuid,
    manufacturer_id uuid,
    name            text,
    categories      text,
    brand           text,
    compound        text,
    gross_mass      numeric,
    net_mass        numeric,
    package_type    text
);

create table if not exists ss.certificate_compliance
(
    id                 uuid primary key,
    product_id         uuid,
    type               text,
    number             text,
    normative_document text,
    status_compliance  boolean,
    registration_date  date,
    expiration_date    date
);

create table if not exists ss.user
(
    id                uuid primary key,
    fio               text,
    phone_number      varchar(30),
    password          text,
    registration_date date
);

create table if not exists ss.promotion
(
    id            uuid primary key,
    type          text,
    description   text,
    discount_size numeric,
    start_date    date,
    end_date      date
);

create table if not exists ss.sale_product
(
    id           uuid primary key,
    shop_id      uuid,
    product_id   uuid,
    promotion_id uuid,
    price        numeric,
    currency     text,
    setting_date date,
    avg_rating   float
);

create table if not exists ss.rating
(
    id              uuid primary key,
    user_id         uuid,
    sale_product_id uuid,
    review          text,
    rating          numeric
);

create table if not exists ss.retailer_distributor
(
    retailer_id    uuid,
    distributor_id uuid,
    primary key (retailer_id, distributor_id)
);

create table if not exists ss.distributor_manufacturer
(
    distributor_id  uuid,
    manufacturer_id uuid,
    primary key (distributor_id, manufacturer_id)
);