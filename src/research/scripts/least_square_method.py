import csv
import math as m

import matplotlib.pyplot as plt
import numpy as np


class LeastSquareMethod:
    """
    Класс аппроксимации табличных значений некоторой простой теоретической функцией
    методом наименьших квадратов
    """

    def __init__(
            self,
            filename: str  # путь к файлу с данными
    ) -> None:
        self.__filename = filename

        self.__arr_x: list[float] = list()
        self.__arr_y: list[float] = list()

        self._n: int = 0

        # self.__read_file()
        self.__read_file_for_index_research()

        plt.figure()

        self.__linear_approx()
        self.__square_approx()
        self.__power_func_approx()
        self.__exponen_func_approx()

    def __read_file(self) -> None:
        """
        Метод парсит файл с датасетом
        :return:
        """
        file = open(file=self.__filename, newline='', encoding='utf-8')

        reader = csv.DictReader(file)

        self._n = 0
        for row in reader:
            self.__arr_x.append(float(row['Requests/s']))
            self.__arr_y.append(float(row['Total Average Response Time']))
            self._n += 1

        file.close()

    def __read_file_for_index_research(self) -> None:
        """
        Метод парсит файл с датасетом
        """
        file = open(file=self.__filename, newline='', encoding='utf-8')

        reader = csv.DictReader(file)

        self._n = 0
        for row in reader:
            self.__arr_x.append(float(row['x']))
            self.__arr_y.append(float(row['without_index']))
            self._n += 1

        file.close()

    def __linear_approx(self) -> None:
        """
        Интерполяция линейной функцией
        :return:
        """

        one = 0
        for x in self.__arr_x:
            one += (x ** 2)

        two = 0
        for x in self.__arr_x:
            two += x

        three = 0
        for i, x in enumerate(self.__arr_x):
            three += (x * self.__arr_y[i])

        four = 0
        for y in self.__arr_y:
            four += y

        mtr_coeff: np.ndarray = np.array([[one, two], [two, self._n]])
        vec_target: np.ndarray = np.array([three, four])

        solve = np.linalg.solve(mtr_coeff, vec_target)

        def approx_func(_x: float):  # полученная аппроксимирующая функция
            return solve[0] * _x + solve[1]

        residual = 0
        for i in range(self._n):
            residual += (approx_func(self.__arr_x[i]) - self.__arr_y[
                i]) ** 2

        print(f"Функция: {solve[0]:.5f} * x + {solve[1]:.5f}")
        print("Невязка при аппроксимации линейной функцией: ", residual)

        x = np.array(self.__arr_x)
        y = np.array(self.__arr_y)

        # Построение графика данных и аппроксимирующей функции
        plt.scatter(x, y, label='Дискретные значения', color='black', marker='o')
        plt.plot(x, approx_func(x), label='Аппроксимирующая линейная функция')
        plt.legend()

        # Включение отображения сетки
        plt.grid(True, linestyle='--', linewidth=0.5, alpha=0.7)
        # Сохранение графика в формате SVG
        plt.savefig(self.__filename[:-3] + 'svg', format='svg')
        plt.axis('equal')

    def __square_approx(self) -> None:
        """
        Интерполяция квадратичной функцией
        :return:
        """

        one = 0
        for x in self.__arr_x:
            one += (x ** 4)

        two = 0
        for x in self.__arr_x:
            two += (x ** 3)

        three = 0
        for x in self.__arr_x:
            three += (x ** 2)

        four = 0
        for x in self.__arr_x:
            four += x

        five = 0
        for i, x in enumerate(self.__arr_x):
            five += ((x ** 2) * self.__arr_y[i])

        six = 0
        for i, x in enumerate(self.__arr_x):
            six += (x * self.__arr_y[i])

        seven = 0
        for y in self.__arr_y:
            seven += y

        mtr_coeff: np.ndarray = np.array([[one, two, three], [two, three, four], [three, four, self._n]])
        vec_target: np.ndarray = np.array([five, six, seven])

        solve = np.linalg.solve(mtr_coeff, vec_target)

        def approx_func(_x: float):  # полученная аппроксимирующая функция
            return solve[0] * (_x ** 2) + solve[1] * _x + solve[2]

        residual = 0
        for i in range(self._n):
            residual += (approx_func(self.__arr_x[i]) - self.__arr_y[
                i]) ** 2

        print(f"Функция: {solve[0]:.5f} * x^2 + {solve[1]:.5f} * x + {solve[2]:.5f}")
        print("Невязка при аппроксимации квадратичной функцией: ", residual)

        x = np.array(self.__arr_x)
        y = np.array(self.__arr_y)

        # Построение графика данных и аппроксимирующей функции
        # plt.scatter(x, y, label='Дискретные значения', color='black', marker='o')
        plt.plot(x, approx_func(x), label='Аппроксимирующая квадратичная функция')
        plt.legend()

        # Включение отображения сетки
        plt.grid(True, linestyle='--', linewidth=0.5, alpha=0.7)
        # Сохранение графика в формате SVG
        plt.savefig(self.__filename[:-3] + 'svg', format='svg')

    def __power_func_approx(self) -> None:
        """
        Интерполяция степенной функцией
        :return:
        """

        one = 0
        for x in self.__arr_x:
            if x != 0:
                one += m.log(x)

        two = 0
        for y in self.__arr_y:
            two += m.log(y)

        three = 0
        for x in self.__arr_x:
            if x != 0:
                three += (m.log(x) ** 2)

        four = 0
        for i in range(self._n):
            if self.__arr_x[i] != 0:
                four += (m.log(self.__arr_x[i]) * m.log(self.__arr_y[i]))

        mtr_coeff: np.ndarray = np.array([[three, one], [one, self._n]])
        vec_target: np.ndarray = np.array([four, two])

        solve = np.linalg.solve(mtr_coeff, vec_target)
        b = m.exp(solve[1])

        def approx_func(_x: np.array):  # полученная аппроксимирующая функция
            return b * (x ** solve[0])

        residual = 0
        for i in range(self._n):
            residual += (approx_func(self.__arr_x[i]) - self.__arr_y[i]) ** 2

        print(f"Функция: {b:.5f} * (x ** {solve[0]:.5f})")
        print("Невязка при аппроксимации степенной функцией: ", residual)

        x = np.array(self.__arr_x)
        y = np.array(self.__arr_y)

        # Построение графика данных и аппроксимирующей функции
        # plt.scatter(x, y, label='Дискретные значения', color='black', marker='o')
        plt.plot(x, approx_func(x), label='Аппроксимирующая степенная функция')
        plt.legend()

        # Включение отображения сетки
        plt.grid(True, linestyle='--', linewidth=0.5, alpha=0.7)
        # Сохранение графика в формате SVG
        plt.savefig(self.__filename[:-3] + 'svg', format='svg')

    def __exponen_func_approx(self) -> None:
        """
        Интерполяция показательной функцией
        :return:
        """
        one = 0
        for x in self.__arr_x:
            one += (x ** 2)

        two = 0
        for y in self.__arr_y:
            two += m.log(y)

        three = 0
        for x in self.__arr_x:
            three += x

        four = 0
        for i in range(self._n):
            four += (self.__arr_x[i] * m.log(self.__arr_y[i]))

        mtr_coeff: np.ndarray = np.array([[one, three], [three, self._n]])
        vec_target: np.ndarray = np.array([four, two])

        solve = np.linalg.solve(mtr_coeff, vec_target)
        b = m.exp(solve[1])

        def approx_func(_x: np.array):  # полученная аппроксимирующая функция
            return b * (np.exp(solve[0] * x))

        residual = 0
        for i in range(self._n):
            residual += (approx_func(self.__arr_x[i]) - self.__arr_y[i]) ** 2

        print(f"Функция: {b:.5f} * (exp({solve[0]:.5f} * x)")
        print("Невязка при аппроксимации показательной функцией: ", residual)

        x = np.array(self.__arr_x)
        y = np.array(self.__arr_y)

        # Построение графика данных и аппроксимирующей функции
        # plt.scatter(x, y, label='Дискретные значения', color='black', marker='o')
        plt.plot(x, approx_func(x), label='Аппроксимирующая показательная функция')
        plt.legend()

        # Включение отображения сетки
        plt.grid(True, linestyle='--', linewidth=0.5, alpha=0.7)
        # Сохранение графика в формате SVG
        plt.savefig(self.__filename[:-3] + 'svg', format='svg')


def main() -> None:
    """
    Главная функция
    :return:
    """

    # LeastSquareMethod('../data/without_cache.csv')
    # LeastSquareMethod('../data/with_cache.csv')
    LeastSquareMethod('../data/index_research1.csv')


if __name__ == "__main__":
    main()
