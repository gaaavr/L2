package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

/* Реализовать утилиту аналог консольной команды cut (man cut).
Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
*/

// Создаём структуру для удобства использования аргументов командной строки
type arguments struct {
	f, d, input string
	s           bool
}

func main() {
	// Задаём и парсим аргументы командной строки
	args, err := setFlags()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Если задан ключ на строки только с разделителем,
	// то в случае отсутствия в строке разделителя не выводим ничего и выходим из программы.
	if args.s {
		if !strings.Contains(args.input, args.d) {
			return
		}
	}
	// Разделяем строку по разделителю на части
	cutStrings := strings.Split(args.input, args.d)
	// Проверяем, что был введён ключ на выбор полей и что строка была разделена,
	// при отрицательном результате колонка будет всего одна
	if args.f != "" && len(cutStrings) > 1 {
		// Получаем массив номеров колонок, которые необходимо вывести
		numCols, err := colSelection(args, cutStrings)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Выводим колонки в соответствии с заданными номерами
		for i, v := range numCols {
			if i == len(numCols)-1 {
				fmt.Print(cutStrings[v] + "\n")
				return
			}
			fmt.Print(cutStrings[v] + args.d)
		}
	}
	// Выводим исходную строку как одну колонку в случае не срабатывания условия проверки строки на разделение
	fmt.Println(args.input)
}

// Функция для  распарсивания аргументов командной строки
func setFlags() (arguments, error) {
	// Объявляем переменные для хранения значений флага
	var (
		f, d string
		s    bool
	)
	// Определяем флаги командной строки и парсим их в переменные
	flag.StringVar(&f, "f", "", "\"fields\" - выбрать поля (колонки)")
	flag.StringVar(&d, "d", "\t", "\"delimiter\" - использовать другой разделитель")
	flag.BoolVar(&s, "s", false, "\"separated\" - только строки с разделителем")
	flag.Parse()
	input := flag.Arg(0)

	// Проверяем на валидное значение введёную строку
	if input == "" {
		return arguments{}, fmt.Errorf("не введена строка")
	}
	// Инициализируем экземпляр структуры с аргументами командной строки и возвращаем их
	args := arguments{
		f:     f,
		d:     d,
		s:     s,
		input: input,
	}
	return args, nil
}

// Функция для определения и выбора колонок из аргумента команлной строки
func colSelection(args arguments, rows []string) (res []int, err error) {
	//  Объёявляем переменную для конвертации строкового значения числа в числовое
	var numberCol int
	// Разбиваем строку аргумента на части для конвертации в числовые значения
	numbers := strings.Split(args.f, ",")
	// В цикле проходим по числам в строковм представлении
	for _, v := range numbers {
		// Проверяем число на содержание знаков "-" для идентификации нескольких колонок подряд
		if strings.Contains(v, "-") && len(v) > 1 && strings.Count(v, "-") == 1 {
			// Если "-" находится перед числом, значит берем все числа от первого и до переданного,
			// конвертируем их в числовое представление и добавляем в массив
			if strings.Index(v, "-") == 0 {
				numberCol, err = strconv.Atoi(strings.TrimPrefix(v, "-"))
				if err != nil {
					return nil, fmt.Errorf("введён неккоректный номер колонки")
				}
				for j := 0; j < numberCol; j++ {
					if j > len(rows)-1 {
						return res, nil
					}
					res = append(res, j)
				}
				continue
				// Если "-" является последнием символом после числа,
				// то берём все числа, пока последнее не станет равно номеру последней колонки,
				// конвертируем их в числовое представление и добавляем в массив
			} else if strings.Index(v, "-") == len(v)-1 {
				numberCol, err = strconv.Atoi(strings.TrimSuffix(v, "-"))
				if err != nil {
					return nil, fmt.Errorf("введён неккоректный номер колонки")
				}
				for j := numberCol - 1; j < len(rows); j++ {
					res = append(res, j)
				}
				continue
				// Если "-" находится в промежутки между двумя числами,
				// берём все числа в этом промежутке и конвертируем их числовое представление,
				// а затем добавляем в массив
			} else {
				twoNumbers := strings.Split(v, "-")
				colOne, err := strconv.Atoi(twoNumbers[0])
				if err != nil {
					return nil, err
				}
				colTwo, err := strconv.Atoi(twoNumbers[1])
				if err != nil {
					return nil, err
				}
				for j := colOne - 1; j < colTwo; j++ {
					if j > len(rows)-1 {
						return res, nil
					}
					res = append(res, j)
				}
				continue
			}
		}
		// Конвертируем строковое представление в числовое, в случае ошибки возвращаем её
		numberCol, err = strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("введён неккоректный номер колонки")
		}
		// Проверяем, что переданное в аргументе число не превышает количество колонок строки
		if numberCol > len(rows) {
			return nil, fmt.Errorf("колонки с номером %d не существует", numberCol)
		}
		// Добавляем в массив полученное число
		res = append(res, numberCol-1)
	}
	// Возвращаем массив номеров колонок, которые необходимо вывести
	return res, nil
}
