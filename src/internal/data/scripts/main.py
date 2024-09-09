from faker import Faker

from generator import Generator

COUNT_RECORDS = 1000

"""
Поменять в схемах в сущности sale_product rating -> avg_rating !!!!!!
"""


def main() -> None:
    faker = Faker("ru_RU")

    generator = Generator(faker=faker)

    generator.retailers_to_csv(COUNT_RECORDS)
    generator.distributors_to_csv(COUNT_RECORDS)
    generator.manufacturers_to_csv(COUNT_RECORDS)
    generator.shops_to_csv()
    generator.products_to_csv(COUNT_RECORDS)
    generator.certificates_compliance_to_csv(1000)
    generator.promotions_to_csv(COUNT_RECORDS)
    generator.sale_products_to_csv(COUNT_RECORDS * COUNT_RECORDS)
    generator.retailer_distributor_to_csv(COUNT_RECORDS)
    generator.distributor_manufacturer_to_csv(COUNT_RECORDS)


if __name__ == '__main__':
    main()
