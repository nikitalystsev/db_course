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

    # def on_start(self) -> None:
    #     sign_in_data = {
    #         "phone_number": "89314022581",
    #         "password": "zhpiix6900"
    #     }
    #     response = self.client.post("/auth/sign-in", json=sign_in_data)
    #     access_token = dict(response.json())['access_token']
    #
    #     self.client.headers.update({"Authorization": f"Bearer {access_token}"})
    # '/products?[limit, offset]'
    # '/sales?[product_id]'
    # '/certificates?[product_id]'
    name = 'test_requests'

    @task(1)
    def get_product(self):
        """
        Получить информацию о товаре
        """
        product_id = random.choice(self.product_ids)
        self.client.get(f"/products/{product_id}", name=self.name)

    @task(1)
    def get_page_products(self):
        """
        Получить страницу товаров
        """
        offset = random.randint(0, 500)
        self.client.get(f"/products?limit=10&offset={offset}", name=self.name)

    @task(1)
    def get_sales_by_product_id(self):
        """
        Запросы на получение продаж товара
        """
        product_id = random.choice(self.product_ids)
        self.client.get(f"/sales?product_id={product_id}", name=self.name)

    @task(1)
    def get_certificates_by_product_id(self):
        """
        Запросы на получение сертификатов по id товара
        """
        product_id = random.choice(self.product_ids)
        self.client.get(f"/certificates?product_id={product_id}", name=self.name)

    # @task(1)
    # def get_sales_by_shop_id(self):
    #     """
    #     Запросы на получение продаж в магазине
    #     """
    #     shop_id = random.choice(self.shop_ids)
    #     self.client.get(f"/api/sales?shop_id={shop_id}", name='/sales?[shop_id]')
