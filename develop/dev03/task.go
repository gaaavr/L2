package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/* Отсортировать строки в файле по аналогии с консольной утилитой sort
(man sort — смотрим описание и основные параметры): на входе подается файл из
несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок,
по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительно

Реализовать поддержку утилитой следующих ключей:

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учетом суффиксов

*/

// Создаём структуру для удобства использования аргументов командной строки
type arguments struct {
	k                   int
	n, r, u, m, b, c, h bool
	input, output       string
}

func main() {
	// Задаём и парсим аргументы командной строки
	args, err := setFlags()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Открываем и считываем данные из файла, названия которого указывается в командной строке
	data, err := openReadFile(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Проверка и приведение номера сортируемой колонки к нужному индексу
	if args.k < 1 {
		args.k = 0
	} else {
		args.k--
	}

	// Выбор функции сортировки в зависимости от выбранных аргументов
	sortFunc := selectSortFunc(data, args)
	// Выполнение проверки на сортированность исходных данных при выбранном соответствующем аргументе
	if args.c {
		if sort.SliceIsSorted(data, sortFunc) {
			fmt.Println("sorted")
			return
		}
		fmt.Println("unsorted")
		return
	}
	// Сортировка исходных данных c помощью функции, выбранной ранее
	sort.Slice(data, sortFunc)

	// Запись отсортированных данных в файл, указанный в аргументах командной строки
	if err = writeToFile(data, args); err != nil {
		fmt.Println(err)
		return
	}

}

// Функция для  распарсивания аргументов командной строки
func setFlags() (arguments, error) {
	// Объявляем переменные для хранения значений флага
	var (
		n, r, u, m, b, c, h bool
		k                   int
	)
	// Определяем флаги командной строки и парсим их в переменные
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющееся строки")
	flag.BoolVar(&m, "m", false, "сортировать по названию месяца")
	flag.BoolVar(&b, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&c, "c", false, "проверять отсортированы ли данные")
	flag.BoolVar(&h, "h", false, "сортировать по числовому значению с учетом суффиксов")
	flag.IntVar(&k, "k", 0, "указание колонки для сортировки")
	flag.Parse()
	input := flag.Arg(0)
	output := flag.Arg(1)

	// Проверяем на валидные имена для файлов чтения и записи данных
	if input == "" || output == "" {
		return arguments{}, fmt.Errorf("не все имена файлов указаны")
	}
	// Инициализируем экземпляр структуры с аргументами командной строки и возвращаем их
	args := arguments{
		k:      k,
		n:      n,
		r:      r,
		u:      u,
		m:      m,
		b:      b,
		c:      c,
		h:      h,
		input:  input,
		output: output,
	}
	return args, nil
}

// Функция чтения данных для сортировки
func openReadFile(args arguments) ([][]string, error) {
	// Открываем файл, имя которого было передано в аргументах
	file, err := os.Open(args.input)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// Создаём сканер для построчного считывания данных из файла
	scanner := bufio.NewScanner(file)
	// Инициализируем матрицу, в которой будут храниться считанные данные
	matrix := make([][]string, 0)
	// Считываем построчно данные
	for scanner.Scan() {
		line := scanner.Text()
		// Обрезаем у строки хвостовые пробелы, если был передан соответствующий аргумент
		if args.b {
			line = strings.TrimSuffix(line, " ")
		}
		// Разбиваем строку на слайс строк и добавляем в матрицу
		words := strings.Split(line, " ")
		matrix = append(matrix, words)
	}
	return matrix, nil
}

// Функция записи отсортированных данных в файл
func writeToFile(data [][]string, args arguments) error {
	// Создаём или обрезаем файл записи
	file, err := os.Create(args.output)
	if err != nil {
		return err
	}
	defer file.Close()
	// Инициализируем мапу для проверки на повторение строк
	repeatingStrings := make(map[string]struct{}, len(data))
	// Инциализируем слайс строк для добавления туда отсортированных строк
	rows := make([]string, 0, len(data))
	for _, v := range data {
		line := strings.Join(v, " ")
		// Проверяем на повторение строки, если был передан соответствующий аргумент
		if args.u {
			// Если строка повторяется, переходим к следущей
			if _, ok := repeatingStrings[line]; ok {
				continue
			}
		}
		repeatingStrings[line] = struct{}{}
		// Добавляем отсортированную строку в слайс
		rows = append(rows, line)
	}
	// Записываем полученные отсортированные данные в файл
	_, err = file.WriteString(strings.Join(rows, "\n"))
	if err != nil {
		return err
	}
	return nil
}

// Функция выбора алгоритма сортировки данных
func selectSortFunc(data [][]string, args arguments) func(i, j int) bool {
	// Объявляем функцию для сортировки данных
	var sortFunc func(i, j int) bool
	// В зависимости от переданных аргументов присваиваем ей алгоритм сравнивания элементов
	switch {
	case args.n:
		// Функция сравнивает данные по числовым значениям
		sortFunc = func(i, j int) bool {
			firstElem, _ := strconv.ParseFloat(getElement(data, i, args.k), 64)
			secElem, _ := strconv.ParseFloat(getElement(data, j, args.k), 64)
			if args.r {
				return firstElem > secElem
			}
			return firstElem < secElem
		}
		// Функция сравнивает данные по названию месяца
	case args.m:
		sortFunc = func(i, j int) bool {
			firstElem, _ := getMonth(getElement(data, i, args.k))
			secElem, _ := getMonth(getElement(data, j, args.k))
			if args.r {
				return firstElem > secElem
			}
			return firstElem < secElem
		}
		// Функция сравнивает данные по суффиксам строк, если номер переданной колонки для сортировки
		// превышает длину строки, то размер элемента равен нулю
	case args.h:
		var firstElem, secElem int
		sortFunc = func(i, j int) bool {
			if args.k >= len(data[i]) {
				firstElem = 0
			} else {
				firstElem = getLen(data[i][args.k:])
			}
			if args.k >= len(data[j]) {
				secElem = 0
			} else {
				secElem = getLen(data[j][args.k:])
			}
			if args.r {
				return firstElem > secElem
			}
			return firstElem < secElem
		}
		// Функция сравнивает данные по размеру
	default:
		sortFunc = func(i, j int) bool {
			firstElem := getElement(data, i, args.k)
			secElem := getElement(data, j, args.k)
			if args.r {
				return firstElem > secElem
			}
			return firstElem < secElem
		}
	}
	return sortFunc
}

// Функция получает элемент исходной строки в зависимости от заданной колонки
func getElement(data [][]string, i, k int) string {
	if k < len(data[i]) {
		return data[i][k]
	}
	return ""
}

// Функция получает номер месяца
func getMonth(month string) (time.Month, error) {
	if m, err := time.Parse("January", month); err == nil {
		return m.Month(), nil
	}
	if m, err := time.Parse("Jan", month); err == nil {
		return m.Month(), nil
	}
	if m, err := time.Parse("01", month); err == nil {
		return m.Month(), nil
	}
	if m, err := time.Parse("1", month); err == nil {
		return m.Month(), nil
	}
	return 0, fmt.Errorf("не удалось получить месяц")
}

// Функция получает длину суффикса строки
func getLen(str []string) int {
	var sumLen int
	for _, v := range str {
		sumLen += len(v)
	}
	return sumLen
}
