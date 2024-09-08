import csv

from matplotlib import pyplot as plt


def parse_stats_history(filepath: str, dirpath: str):
    """
    Парсит файл с историей статистики для каждого типа запроса по отдельности
    """
    stats_file = open(file=filepath, newline='', encoding='utf-8')
    get_product_file = open(file=dirpath + "get_product.csv", mode='w', newline='', encoding='utf-8')
    get_page_product_file = open(file=dirpath + "get_page_product.csv", mode='w', newline='',
                                 encoding='utf-8')
    get_sales_by_product_id_file = open(file=dirpath + "get_sales_by_product_id.csv", mode='w',
                                        newline='', encoding='utf-8')
    get_certificates_by_product_id_file = open(file=dirpath + "get_certificates_by_product_id.csv",
                                               mode='w', newline='', encoding='utf-8')
    fieldnames = [
        'Total Average Response Time',
        'Requests/s'
    ]
    reader = csv.DictReader(stats_file)
    get_product_writer = csv.DictWriter(get_product_file, fieldnames=fieldnames)
    get_product_writer.writeheader()
    get_page_product_writer = csv.DictWriter(get_page_product_file, fieldnames=fieldnames)
    get_page_product_writer.writeheader()
    get_sales_by_product_id_writer = csv.DictWriter(get_sales_by_product_id_file, fieldnames=fieldnames)
    get_sales_by_product_id_writer.writeheader()
    get_certificates_by_product_id_writer = csv.DictWriter(get_certificates_by_product_id_file, fieldnames=fieldnames)
    get_certificates_by_product_id_writer.writeheader()

    for row in reader:
        match row['Name']:
            case '/certificates?product_id':
                get_certificates_by_product_id_writer.writerow({
                    'Total Average Response Time': row['Total Average Response Time'],
                    'Requests/s': row['Requests/s'],
                })
            case '/products/id':
                get_product_writer.writerow({
                    'Total Average Response Time': row['Total Average Response Time'],
                    'Requests/s': row['Requests/s'],
                })
            case '/products?limit&offset':
                get_page_product_writer.writerow({
                    'Total Average Response Time': row['Total Average Response Time'],
                    'Requests/s': row['Requests/s'],
                })
            case '/sales?product_id':
                get_sales_by_product_id_writer.writerow({
                    'Total Average Response Time': row['Total Average Response Time'],
                    'Requests/s': row['Requests/s'],
                })
            case _:
                continue

    get_certificates_by_product_id_file.close()
    get_sales_by_product_id_file.close()
    get_page_product_file.close()
    get_product_file.close()
    stats_file.close()


def build_graphics(filepath1: str, filepath2: str, output_svg: str):
    without_cache_file = open(file=filepath1, newline='', encoding='utf-8')
    with_cache_file = open(file=filepath2, newline='', encoding='utf-8')
    without_cache_reader = csv.DictReader(without_cache_file)
    with_cache_reader = csv.DictReader(with_cache_file)

    without_cache_data = [[], []]
    for row in without_cache_reader:
        without_cache_data[0].append(float(row['Total Average Response Time']))
        without_cache_data[1].append(float(row['Requests/s']))

    with_cache_data = [[], []]
    for row in with_cache_reader:
        with_cache_data[0].append(float(row['Total Average Response Time']))
        with_cache_data[1].append(float(row['Requests/s']))

    plt.figure()
    plt.plot(without_cache_data[1], without_cache_data[0], label='Без использования кеша', marker='x')
    plt.plot(with_cache_data[1], with_cache_data[0], label='С использованием кеша', marker='*')

    plt.title('Исследование')
    plt.ylabel('Среднее время ответа')
    plt.xlabel('Число запросов в секунду')

    plt.legend()
    plt.grid()
    # plt.show()

    plt.savefig(output_svg, format='svg')

    with_cache_file.close()
    without_cache_file.close()


def main():
    parse_stats_history(
        "../locust_stats/without_cache_stats_history.csv",
        "../data/without_cache/"
    )
    parse_stats_history(
        "../locust_stats/with_cache_stats_history.csv",
        "../data/with_cache/"
    )

    build_graphics(
        "../data/without_cache/get_product.csv",
        "../data/with_cache/get_product.csv",
        "../graphics/get_product.svg"
    )

    build_graphics(
        "../data/without_cache/get_page_product.csv",
        "../data/with_cache/get_page_product.csv",
        "../graphics/get_page_product.svg"
    )

    build_graphics(
        "../data/without_cache/get_sales_by_product_id.csv",
        "../data/with_cache/get_sales_by_product_id.csv",
        "../graphics/get_sales_by_product_id.svg"
    )

    build_graphics(
        "../data/without_cache/get_certificates_by_product_id.csv",
        "../data/with_cache/get_certificates_by_product_id.csv",
        "../graphics/get_certificates_by_product_id.svg"
    )


if __name__ == '__main__':
    main()
