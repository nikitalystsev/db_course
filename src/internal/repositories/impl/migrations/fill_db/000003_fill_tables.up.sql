copy ss.retailer from '/data/retailer.csv' with delimiter ',' csv header;

copy ss.distributor from '/data/distributor.csv' with delimiter ',' csv header;

copy ss.manufacturer from '/data/manufacturer.csv' with delimiter ',' csv header;

copy ss.shop from '/data/shop.csv' with delimiter ',' csv header;

copy ss.product from '/data/product.csv' with delimiter ',' csv header;

copy ss.certificate_compliance from '/data/certificate_compliance.csv' with delimiter ',' csv header;

copy ss.promotion from '/data/promotion.csv' with delimiter ',' csv header;

copy ss.sale_product from '/data/sale_product.csv' with delimiter ',' csv header;

copy ss.retailer_distributor from '/data/retailer_distributor.csv' with delimiter ',' csv header;

copy ss.distributor_manufacturer from '/data/distributor_manufacturer.csv' with delimiter ',' csv header;

insert into ss."user"
values ('75919792-c2d9-4685-92b2-e2a80b2ed5be', 'Randall C. Jernigan', '79314562376',
        '$2a$10$8APnhcfxoGxXGdNHSdBEaebuwcIkjwEnSHOIv.xu9bmROkpCRLTJS', CURRENT_DATE), -- пароль: sdgdgsgsgd
       ('5818061a-662d-45bb-a67c-0d2873038e65', 'Jesse M. Flores', '72443564633',
        '$2a$10$2cYeMgl8fjH76HjIm54enOuHUiV3qzV81jdVJLLNCQbo2zXc9jija', CURRENT_DATE), -- пароль: qwresdfdsf
       ('3885b2d3-ef6e-4f62-8f86-d1454d108207', 'Mitrofan Bogdanov', '76867456521',
        '$2a$10$KxEprnJtxnL./4Zts.IP3uOQfGktXZXTp1BmvMKZxyDSJoIm4hmt6', CURRENT_DATE), -- пароль: hghhfnnbdd
       ('6800b3ee-9810-450e-9ca5-776aa1c6191d', 'Peter Zuev', '32534523451',
        '$2a$10$GjKIYnr6wRohYWkUhmlPhO5uza1zvudS9rWeydAv1yzEW0GfTOAme', CURRENT_DATE), -- пароль: rtjhhhgffr
       ('8d9b001f-5760-4c40-bc60-988e0ca54d18', 'Vasilisa Agapova', '73453562423',
        '$2a$10$sQZzp5BlhAvTMc/AIzAUS.PVuAxxH/rVmNfv.W73RhdxH7xSdbyQy', CURRENT_DATE), -- пароль: gfjkjdgffy
       ('362b79f6-d671-404a-b1a0-5a655aebc1b6', 'Лысцев Никита Дмитриевич', '89314022581',
        '$2a$10$xDzRFS0ClhEcosyFVQEPCev8AXakZyYau4Hk8iN3dyTXJYXUj1coO', CURRENT_DATE);