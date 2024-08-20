import csv
import datetime
import random
import string
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
    'ТР ТС 029/2012 «Требования безопасности пищевых добавок, '
    'ароматизаторов и технологических вспомогательных средств»',
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
                certificate_type = random.choice(CERTIFICATE_COMPLIANCE_TYPES)

                writer.writerow({
                    "id": certificate_compliance_id,
                    "product_id": product_id,
                    "type": certificate_type,
                    "number": self.__get_certificate_number_by_type(certificate_type),
                    "normative_document": random.choice(CERTIFICATE_COMPLIANCE_NORMATIVE_DOCUMENTS),
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

    @staticmethod
    def __get_vsd_number():
        """
        Генерирует номер ветеринарной сопроводительной документации в формате
        A653-D6D8-9B92-4C05-BFDA-513D-80B3-E48D.
        """
        segments = []

        for _ in range(8):  # Всего 8 сегментов
            # Генерируем сегмент: 4 символа (буквы и цифры) + 4 цифры
            segment = ''.join(random.choices(string.ascii_uppercase + string.digits, k=4))
            segment += '-' + ''.join(random.choices(string.digits, k=4))
            segments.append(segment)

        # Объединяем сегменты в одну строку
        document_number = '-'.join(segments)

        return document_number

    @staticmethod
    def __get_declaration_number():
        """
        Генерирует номер декларации о соответствии в формате:
        ТС N Д-RU С-NL.НВ96.В.01216/24
        """

        # Генерируем случайные части номера
        ts_number = "ТС N Д-" + random.choice(["RU", "KZ", "BY"])  # Пример: RU, KZ, BY
        country_code = random.choice(["RU", "NL", "US", "CN"])  # Пример стран
        nv_part = "НВ" + str(random.randint(10, 99))  # Номер от 10 до 99
        version_part = "В." + str(random.randint(1, 9))  # Версия от 1 до 9
        serial_number = str(random.randint(10000, 99999))  # Серийный номер от 10000 до 99999
        year_part = str(random.randint(20, 24))  # Год от 20 до 24

        # Формируем итоговый номер декларации
        declaration_number = f"{ts_number} С-{country_code}.{nv_part}.{version_part}.{serial_number}/{year_part}"

        return declaration_number

    @staticmethod
    def __get_gost_certificate_number():
        """
        Генерирует номер сертификата соответствия ГОСТ Р в формате:
        РОСС RU.0001.11МТ49
        """

        # Генерируем случайные части номера
        prefix = "РОСС"
        country_code = "RU"
        registration_number = f"{random.randint(1000, 9999)}"  # Четырехзначный регистрационный номер
        year_code = f"{random.randint(10, 99)}"  # Двухзначный код года
        suffix = ''.join(random.choices('ABCDEFGHIJKLMNOPQRSTUVWXYZ', k=2)) + str(
            random.randint(10, 99))  # Две буквы и двухзначное число

        # Формируем итоговый номер сертификата
        certificate_number = f"{prefix} {country_code}.{registration_number}.{year_code}{suffix}"

        return certificate_number

    @staticmethod
    def __get_sgr_number():
        """
        Генерирует номер СГР на пищевую продукцию в формате:
        RU.77.99.11.003.E.021866.05.11
        """

        # Генерируем случайные части номера
        country_code = "RU"
        region_code = f"{random.randint(10, 99)}"  # Двухзначный код региона
        category_code = f"{random.randint(10, 99)}"  # Двухзначный код категории
        product_code = f"{random.randint(10, 99)}"  # Двухзначный код продукта
        registration_number = f"{random.randint(100, 999)}"  # Трехзначный регистрационный номер
        status_code = random.choice(['E', 'C'])  # Код статуса (E или C)
        serial_number = f"{random.randint(100000, 999999)}"  # Шестизначный серийный номер
        date_code = f"{random.randint(1, 12):02}.{random.randint(1, 31):02}"  # Дата (месяц.число)

        # Формируем итоговый номер СГР
        sgr_number = (f"{country_code}.{region_code}.{category_code}."
                      f"{product_code}.{registration_number}.{status_code}.{serial_number}.{date_code}")

        return sgr_number

    def __get_certificate_number_by_type(self, certificate_type: str):
        """
        Метод генерирует номер сертификата по его типу
        """
        if certificate_type == CERTIFICATE_COMPLIANCE_TYPES[1]:
            return self.__get_declaration_number()

        if certificate_type == CERTIFICATE_COMPLIANCE_TYPES[2]:
            return self.__get_gost_certificate_number()

        if certificate_type == CERTIFICATE_COMPLIANCE_TYPES[3]:
            return self.__get_sgr_number()

        return self.__get_vsd_number()
