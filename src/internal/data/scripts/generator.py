import csv
import datetime
import random
import uuid

from faker import Faker
from mimesis import Food
from mimesis.locales import Locale

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

PRODUCT_CATEGORIES = [
    'говядина, баранина, свинина (кроме бескостного мяса)',
    'куры (кроме куриных окорочков)',
    'рыба мороженая неразделанная',
    'масло сливочное',
    'яйца куриные',
    'сахар-песок',
    'соль поваренная пищевая',
    'мука пшеничная',
    'хлеб ржаной, ржано-пшеничный',
    'хлеб и булочные изделия из пшеничной муки',
    'рис шлифованный',
    'пшено',
    'крупа гречневая – ядрица',
    'вермишель',

    'картофель',
    'капуста белокочанная свежая',
    'лук репчатый',
    'морковь',

    'яблоки'

    'масло подсолнечное',
    'молоко питьевое',
    'чай черный байховый'
]

PRODUCT_PACKAGE_TYPES = [
    'Пластиковая упаковка',
    'Картонная коробка',
    'Стеклянная банка',
    'Полиэтиленовый пакет'
]

CERTIFICATE_COMPLIANCE_TYPES = [
    'Транспортная ВСД',
    'декларация о соответствии на пищевую продукцию (ДС)',
    'добровольный сертификат на пищевую продукцию',
    'СГР на пищевую продукцию'
]

CERTIFICATE_COMPLIANCE_NORMATIVE_DOCUMENTS = [
    'ТР ТС 021/2011 «О безопасности пищевой продукции»',
    'ТР ТС 022/2011 «Пищевая продукция в части ее маркировки»',
    'ТР ТС 015/2011 «О безопасности зерна»',
    'ТР ТС 023/2011 «Технический регламент на соковую продукцию из фруктов и овощей»',
    'ТР ТС 024/2011 «Технический регламент на масложировую продукцию»',
    'ТР ТС 027/2012 «О безопасности отдельных видов специализированной пищевой '
    'продукции, в том числе диетического лечебного и диетического профилактического питания»',
    'ТР ТС 029/2012 «Требования безопасности пищевых добавок, ароматизаторов и технологических вспомогательных средств»',
    'ТР ТС 033/2013 «О безопасности молока и молочной продукции»',
    'ТР ТС 034/2013 «О безопасности мяса и мясной продукции»',
    'ТР ЕАЭС 040/2016 «О безопасности рыбы и рыбной продукции»'
]


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

    def products_to_csv(self, num: int):
        """
        Метод для генерации продуктов
        """
        if not self.retailer_ids or not self.distributor_ids or not self.manufacturer_ids:
            print("Нет либо ритейлеров, либо дистрибьюторов, либо производителей")
            return

        self.product_retailer_ids.clear()

        with open(file=PRODUCT_DATA_PATH, mode='w', newline='', encoding='utf-8') as file:
            writer = csv.DictWriter(
                file,
                fieldnames=[
                    "id",
                    "retailer_id",
                    "distributor_id",
                    "manufacturer_id",
                    "name",
                    "categories",
                    "brand",
                    "compound",
                    "gross_mass",
                    "net_mass",
                    "package_type"
                ]
            )
            writer.writeheader()

            food = Food(locale=Locale.RU)

            brand_faker = Faker()

            for _ in range(num):
                product_id = str(uuid.uuid4())
                retailer_id = random.choice(self.retailer_ids)
                category = random.choice(PRODUCT_CATEGORIES)
                writer.writerow({
                    "id": product_id,
                    "retailer_id": retailer_id,
                    "distributor_id": random.choice(self.distributor_ids),
                    "manufacturer_id": random.choice(self.manufacturer_ids),
                    "name": self.__get_product_name_by_category(category, food),
                    "categories": category,
                    "brand": f"{brand_faker.word().capitalize()}{brand_faker.word().capitalize()}",
                    "compound": ", ".join([food.spices() for _ in range(random.randint(3, 5))]),
                    "gross_mass": round(random.uniform(0.3, 5), 2),
                    "net_mass": round(random.uniform(0.1, 4.9), 2),
                    "package_type": random.choice(PRODUCT_PACKAGE_TYPES)
                })
                self.product_retailer_ids.append((product_id, retailer_id))

    def certificates_compliance_to_csv(self, num: int):
        """
        Метод для генерации продуктов
        """
        if not self.product_retailer_ids:
            print("Нет товаров")
            return

        with open(file=PRODUCT_DATA_PATH, mode='w', newline='', encoding='utf-8') as file:
            writer = csv.DictWriter(
                file,
                fieldnames=[
                    "id",
                    "product_id",
                    "type",
                    "number",
                    "normative_document",
                    "status_compliance",
                    "registration_data",
                    "expiration_data"
                ]
            )
            writer.writeheader()
            for _ in range(num):
                certificate_compliance_id = str(uuid.uuid4())
                product_id = random.choice(self.product_retailer_ids)[0]
                dates = self.__get_random_dates()

                writer.writerow({
                    "id": certificate_compliance_id,
                    "product_id": product_id,
                    "status_compliance": True,
                    "registration_data": dates[0],
                    "expiration_data": dates[1]
                })

    @staticmethod
    def __get_product_name_by_category(category: str, food_faker: Food):
        """
        Метод для генерации имени товара по его категории
        """
        if category in PRODUCT_CATEGORIES[14:18]:
            return food_faker.vegetable()

        if category == PRODUCT_CATEGORIES[18]:
            return food_faker.fruit()

        if category in PRODUCT_CATEGORIES[19:]:
            return food_faker.drink()

        return food_faker.dish()

    @staticmethod
    def __get_random_dates():
        # Получаем текущую дату и время
        now = datetime.datetime.now()

        # Генерируем случайное количество дней для первой даты (от 0 до 30)
        days_delta_1 = random.randint(0, 365)
        date_1 = now - datetime.timedelta(days=days_delta_1)

        # Генерируем случайное количество дней для второй даты (от 0 до 30, но больше первой)
        days_delta_2 = random.randint(days_delta_1 + 1, days_delta_1 + 365)
        date_2 = now - datetime.timedelta(days=days_delta_2)

        return date_1, date_2
