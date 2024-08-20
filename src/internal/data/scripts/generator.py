import csv
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

    def retailers_to_csv(self, num: int):
        """
        Метод для генерации ритейлеров
        """
        with open(file=RETAILER_DATA_PATH, mode='w', newline='', encoding='utf-8') as file:
            writer = csv.DictWriter(file, fieldnames=["id", "name", "address", "phone_number", "fio_representative"])
            writer.writeheader()

            for _ in range(num):
                writer.writerow({
                    "id": str(uuid.uuid4()),
                    "name": self.faker.company(),
                    "address": self.faker.address().replace("\n", ", "),
                    "phone_number": self.faker.phone_number(),
                    "fio_representative": self.faker.name()
                })

    def distributors_to_csv(self, num: int):
        """
        Метод для генерации дистрибьюторов
        """
        with open(file=DISTRIBUTOR_DATA_PATH, mode='w', newline='', encoding='utf-8') as file:
            writer = csv.DictWriter(file, fieldnames=["id", "name", "address", "phone_number", "fio_representative"])
            writer.writeheader()

            for _ in range(num):
                writer.writerow({
                    "id": str(uuid.uuid4()),
                    "name": self.faker.company(),
                    "address": self.faker.address().replace("\n", ", "),
                    "phone_number": self.faker.phone_number(),
                    "fio_representative": self.faker.name()
                })

    def manufacturers_to_csv(self, num: int):
        """
        Метод для генерации производителей
        """
        with open(file=MANUFACTURER_DATA_PATH, mode='w', newline='', encoding='utf-8') as file:
            writer = csv.DictWriter(file, fieldnames=["id", "name", "address", "phone_number", "fio_representative"])
            writer.writeheader()

            for _ in range(num):
                writer.writerow({
                    "id": str(uuid.uuid4()),
                    "name": self.faker.company(),
                    "address": self.faker.address().replace("\n", ", "),
                    "phone_number": self.faker.phone_number(),
                    "fio_representative": self.faker.name()
                })
