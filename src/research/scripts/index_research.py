# Функция для парсинга данных из файла
import csv

from matplotlib import pyplot as plt


def parse_data(file_path):
    n_values = []
    time1_values = []
    time2_values = []

    with open(file_path, 'r') as file:
        for line in file:
            if '\\hline' in line:  # Игнорируем строки с \hline
                continue
            # Убираем символы \\ и разделяем по &
            parts = line.strip().replace('\\', '').split('&')
            if len(parts) == 3:
                # Приводим данные к нужным типам
                n_values.append(int(parts[0].strip()))
                time1_values.append(float(parts[1].strip()))
                time2_values.append(float(parts[2].strip()))

    return n_values, time1_values, time2_values


def write_to_csv(x, without_index, with_index, output_file):
    """
    Функция для записи данных в CSV файл
    """
    with open(output_file, mode='w', newline='') as file:
        writer = csv.DictWriter(file, fieldnames=['x', 'without_index', 'with_index'])

        writer.writeheader()

        for n, t1, t2 in zip(x, without_index, with_index):
            writer.writerow({
                'x': n,
                'without_index': t1,
                'with_index': t2,
            })


def build_graphics(
        x, without_index, with_index, output_svg
):
    plt.figure()

    plt.plot(x, without_index, marker='o', label='Без использования индекса')
    plt.plot(x, with_index, marker='s', label='С использованием индекса')

    plt.xlabel('Число записей в таблице')
    plt.ylabel('Время выполнения запроса, мкс')
    plt.legend()
    plt.grid()

    plt.savefig(output_svg, format='svg')


def main():
    # Путь к файлу с данными
    file_path = '../latex/index_research1.txt'
    output_file_path = '../data/index_research1.csv'

    x, without_index, with_index = parse_data(file_path)
    write_to_csv(x, without_index, with_index, output_file_path)

    build_graphics(x, without_index, with_index, "../../../docs/coursework/inc/img/index-research1.svg")


if __name__ == "__main__":
    main()
