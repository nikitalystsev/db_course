import csv
import random

from locust import HttpUser, task, constant


def read_id_column(file_path):
    ids = []
    with open(file_path, mode='r', newline='', encoding='utf-8') as csvfile:
        reader = csv.DictReader(csvfile)
        for row in reader:
            ids.append(row['id'])
    return ids


class SmartShopperTestUser(HttpUser):
    wait_time = constant(0)

    product_ids = read_id_column("../../internal/data/data/product.csv")
    shop_ids = read_id_column("../../internal/data/data/shop.csv")

    @task(1)
    def get_product(self):
        """
        Получить информацию о товаре
        """
        product_id = random.choice(self.product_ids)
        self.client.get(f"/products/{product_id}", name="/products/id")

    @task(1)
    def get_page_products(self):
        """
        Получить страницу товаров
        """
        offset = random.randint(0, 500)
        self.client.get(f"/products?limit=10&offset={offset}", name='/products?limit&offset')

    @task(1)
    def get_sales_by_product_id(self):
        """
        Запросы на получение продаж товара
        """
        product_id = random.choice(self.product_ids)
        self.client.get(f"/sales?product_id={product_id}", name='/sales?product_id')

    @task(1)
    def get_certificates_by_product_id(self):
        """
        Запросы на получение сертификатов по id товара
        """
        product_id = random.choice(self.product_ids)
        self.client.get(f"/certificates?product_id={product_id}", name="/certificates?product_id")
