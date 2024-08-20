import csv
import random
import uuid

from faker import Faker

RETAILER_DATA_PATH = "../data/retailer.csv"
DISTRIBUTOR_DATA_PATH = "../data/distributor.csv"
MANUFACTURER_DATA_PATH = "../data/manufacturer.csv"
SHOP_DATA_PATH = "../data/shop.csv"
PRODUCT_DATA_PATH = "../data/product.csv"
CERTIFICATE_COMPLIANCE_DATA_PATH = "../data/certificate_compliance.csv"
USER_DATA_PATH = "../data/user.csv"
PRICE_DATA_PATH = "../data/price.csv"
PROMOTION_DATA_PATH = "../data/promotion.csv"
SALE_PRODUCT_DATA_PATH = "../data/sale_product.csv"
RATING_DATA_PATH = "../data/rating.csv"
RETAILER_DISTRIBUTOR_DATA_PATH = "../data/retailer_distributor.csv"
DISTRIBUTOR_MANUFACTURER_DATA_PATH = "../data/distributor_manufacturer.csv"

PARSE_SHOP_DATA_PATH = "../parse/shops.csv"


class Generator:
    """
    Класс для генерации данных
    """

    def __init__(self, faker: Faker):
        self.faker = faker

        self.retailer_ids = []
        self.distributor_ids = []
        self.manufacturer_ids = []

        self.shop_retailer_ids = []
        self.product_retailer_ids = []

    def retailers_to_csv(self, num: int):
        """
        Метод для генерации ритейлеров
        """
        self.retailer_ids.clear()

        with open(file=RETAILER_DATA_PATH, mode='w', newline='', encoding='utf-8') as file:
            writer = csv.DictWriter(file, fieldnames=["id", "name", "address", "phone_number", "fio_representative"])
            writer.writeheader()

            for _ in range(num):
                retailer_id = str(uuid.uuid4())
                writer.writerow({
                    "id": retailer_id,
                    "name": self.faker.company(),
                    "address": self.faker.address().replace("\n", ", "),
                    "phone_number": self.faker.phone_number(),
                    "fio_representative": self.faker.name()
                })
                self.retailer_ids.append(retailer_id)

    def distributors_to_csv(self, num: int):
        """
        Метод для генерации дистрибьюторов
        """
        self.distributor_ids.clear()

        with open(file=DISTRIBUTOR_DATA_PATH, mode='w', newline='', encoding='utf-8') as file:
            writer = csv.DictWriter(file, fieldnames=["id", "name", "address", "phone_number", "fio_representative"])
            writer.writeheader()

            for _ in range(num):
                distributor_id = str(uuid.uuid4())
                writer.writerow({
                    "id": distributor_id,
                    "name": self.faker.company(),
                    "address": self.faker.address().replace("\n", ", "),
                    "phone_number": self.faker.phone_number(),
                    "fio_representative": self.faker.name()
                })
                self.distributor_ids.append(distributor_id)

    def manufacturers_to_csv(self, num: int):
        """
        Метод для генерации производителей
        """
        self.manufacturer_ids.clear()

        with open(file=MANUFACTURER_DATA_PATH, mode='w', newline='', encoding='utf-8') as file:
            writer = csv.DictWriter(file, fieldnames=["id", "name", "address", "phone_number", "fio_representative"])
            writer.writeheader()

            for _ in range(num):
                manufacturer_id = str(uuid.uuid4())
                writer.writerow({
                    "id": manufacturer_id,
                    "name": self.faker.company(),
                    "address": self.faker.address().replace("\n", ", "),
                    "phone_number": self.faker.phone_number(),
                    "fio_representative": self.faker.name()
                })
                self.manufacturer_ids.append(manufacturer_id)

    def shops_to_csv(self, num: int):
        """
        Метод для генерации магазинов
        """
        if not self.retailer_ids:
            print("Нет ритейлеров")
            return

        self.shop_retailer_ids.clear()

        parse_file = open(file=PARSE_SHOP_DATA_PATH, mode='r', newline='', encoding='utf-8')
        shop_file = open(file=SHOP_DATA_PATH, mode='w', newline='', encoding='utf-8')

        reader = csv.DictReader(parse_file)

        writer = csv.DictWriter(
            shop_file,
            fieldnames=["id", "retailer_id", "title", "address", "phone_number", "fio_director"]
        )
        writer.writeheader()

        cnt = 1

        for row in reader:
            shop_id = str(uuid.uuid4())
            retailer_id = random.choice(self.retailer_ids)

            address = (row['Страна'] + ',' + row['Регион'] + ',' +
                       row['Населенный Пункт'] + ',' + row['Индекс'] + ',' + row['Адрес'])
            writer.writerow({
                "id": shop_id,
                "retailer_id": retailer_id,
                "title": row['Название'],
                "address": address,
                "phone_number": row['Сотовый'] if row['Сотовый'] != "l" else self.faker.phone_number(),
                "fio_director": self.faker.name()
            })
            self.shop_retailer_ids.append((shop_id, retailer_id))
            cnt += 1

            if cnt >= num:
                break

        shop_file.close()
        parse_file.close()
