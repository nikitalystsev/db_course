import csv

import matplotlib.pyplot as plt


def parse_stats_history(filepath: str, out_filepath: str, latex_out_path: str):
    """
    Функция парсит историю статистики исследования
    """
    stats_file = open(file=filepath, newline='', encoding='utf-8')
    out_file = open(file=out_filepath, mode='w', newline='', encoding='utf-8')
    latex_file = open(file=latex_out_path, mode='w', newline='', encoding='utf-8')

    reader = csv.DictReader(stats_file)

    out_fieldnames = [
        'User Count',
        'Total Request Count',
        'Total Average Response Time',
        'Requests/s',
    ]

    writer = csv.DictWriter(out_file, fieldnames=out_fieldnames)
    writer.writeheader()

    count = 0  # Счетчик для каждой строки

    for row in reader:
        if row['Name'] == 'Aggregated':
            continue

        if count % 2 == 0:
            writer.writerow({
                'User Count': row['User Count'],
                'Total Request Count': row['Total Request Count'],
                'Total Average Response Time': round(float(row['Total Average Response Time']), 3),
                'Requests/s': round(float(row['Requests/s']))
            })

            latex_file.write(
                f"{round(float(row['Requests/s']))} & {float(row['Total Average Response Time']):.3f} \\\\ \n\\hline\n"
            )

        count += 1

    latex_file.close()
    out_file.close()
    stats_file.close()


def build_graphics(
        filepath1: str,
        filepath2: str,
        output_svg: str
):
    """
    Функция строит графики исследований
    """
    without_cache_file = open(file=filepath1, newline='', encoding='utf-8')
    with_cache_file = open(file=filepath2, newline='', encoding='utf-8')

    without_cache_reader = csv.DictReader(without_cache_file)
    with_cache_reader = csv.DictReader(with_cache_file)

    without_cache_data = [[], []]
    with_cache_data = [[], []]

    for row in without_cache_reader:
        without_cache_data[0].append(float(row['Total Average Response Time']))
        without_cache_data[1].append(float(row['Requests/s']))

    for row in with_cache_reader:
        with_cache_data[0].append(float(row['Total Average Response Time']))
        with_cache_data[1].append(float(row['Requests/s']))

    plt.plot(without_cache_data[1], without_cache_data[0], label='Без использования кеширования', color='blue', marker='x')
    plt.plot(with_cache_data[1], with_cache_data[0], label='С использованием кеширования', color='red', marker='*')

    plt.ylabel('Среднее время ответа, мс')
    plt.xlabel('Число запросов в секунду')

    plt.legend()
    plt.grid()
    # plt.show()

    # Устанавливаем одинаковый масштаб для осей
    plt.axis('equal')

    plt.savefig(output_svg, format='svg')

    with_cache_file.close()
    without_cache_file.close()


def main() -> None:
    parse_stats_history(
        '../locust_stats/without_cache_stats_history.csv',
        '../data/without_cache.csv',
        '../latex/without_cache.txt',
    )

    parse_stats_history(
        '../locust_stats/with_cache_stats_history.csv',
        '../data/with_cache.csv',
        '../latex/with_cache.txt',
    )

    build_graphics(
        '../data/without_cache.csv',
        '../data/with_cache.csv',
        "../../../docs/coursework/inc/img/research.svg"
    )


if __name__ == '__main__':
    main()

# locust --host=http://localhost:8000 --headless --csv=../locust_stats/without_cache -u 500 -r 10 -t 1m
