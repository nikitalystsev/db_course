from faker import Faker

from generator import Generator

COUNT_RECORDS = 1000


def main() -> None:
    faker = Faker("ru_RU")

    generator = Generator(faker=faker)
    generator.retailers_to_csv(COUNT_RECORDS)
    generator.distributors_to_csv(COUNT_RECORDS)
    generator.manufacturers_to_csv(COUNT_RECORDS)
    generator.shops_to_csv(COUNT_RECORDS)
    generator.products_to_csv(COUNT_RECORDS)
    generator.certificates_compliance_to_csv(COUNT_RECORDS)
    generator.prices_to_csv(COUNT_RECORDS)
    generator.promotions_to_csv(COUNT_RECORDS)


if __name__ == '__main__':
    main()
