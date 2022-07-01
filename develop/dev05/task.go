package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).
Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", напечатать номер строки
*/

// Структура для удобства представления аргументов командной строки
type arguments struct {
	A, B, C, idx           int
	c, i, v, F, n          bool
	oldMatch, match, input string
}

func main() {
	// Задаём и парсим аргументы командной строки
	args, err := setFlags()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Считываем данные из файла
	rows, err := openReadFile(&args)
	if err != nil {
		fmt.Println(err)
		return
	}
	if args.F {
		// Ищет позицию полной совпадающий строки
		findFullString(&args, rows)
	} else {
		// Ищет первое вхождение переданной строки
		findPosString(&args, rows)
	}
	if args.n {
		// При совпадении строк выводится позиция вхождения
		if args.idx >= 0 {
			fmt.Printf("-n:\n\tПозиция строки '%s': %d\n", args.oldMatch, args.idx+1)
		} else {
			fmt.Printf("-n:\n\tСтрока '%s' не найдена\n", args.oldMatch)
		}
	}
	if args.c {
		// Подсчитывается количество вхождений в строки совпадающих слов
		count := countString(args, rows)
		fmt.Printf("-c:\n\tКоличество строк, содержащих '%s': %d\n", args.oldMatch, count)
	}

	if args.v {
		// Удаляем совпадающие элементы из строк, если они есть и завершаем программу
		if args.idx >= 0 {
			removeString(args, rows)
			return
		} else {
			fmt.Printf("-v:\n\tСтрока '%s' не найдена\n", args.oldMatch)
			return
		}
	}
	if args.A > 0 {
		findStringAfter(args, rows)
	}
	if args.B > 0 {
		findStringBefore(args, rows)
	}
}

// Функция для  распарсивания аргументов командной строки
func setFlags() (arguments, error) {
	// Объявляем переменные для хранения значений флага
	var (
		count, i, v, f, n bool
		a, b, c           int
	)
	// Определяем флаги командной строки и парсим их в переменные
	flag.BoolVar(&count, "c", false, "\"count\" (количество строк)")
	flag.BoolVar(&i, "i", false, "\"ignore-case\" (игнорировать регистр)")
	flag.BoolVar(&v, "v", false, "\"invert\" (вместо совпадения, исключать)")
	flag.BoolVar(&f, "F", false, "\"fixed\", точное совпадение со строкой, не паттерн")
	flag.BoolVar(&n, "n", false, "\"line num\", напечатать номер строки")
	flag.IntVar(&a, "A", 0, "\"after\" печатать +N строк после совпадения")
	flag.IntVar(&b, "B", 0, "\"before\" печатать +N строк до совпадения")
	flag.IntVar(&c, "C", 0, "\"context\" (A+B) печатать ±N строк вокруг совпадения")
	flag.Parse()
	match := flag.Arg(0)
	input := flag.Arg(1)

	// Проверяем на валидные имена для файлов чтения данных
	if input == "" || match == "" {
		return arguments{}, fmt.Errorf("не указано имя файла для чтения или искомое слово")
	}
	// Инициализируем экземпляр структуры с аргументами командной строки, а также индексом вхождения
	// искомой строки и возвращаем их
	args := arguments{
		A:        a,
		B:        b,
		C:        c,
		c:        count,
		i:        i,
		v:        v,
		F:        f,
		n:        n,
		oldMatch: match,
		match:    match,
		input:    input,
		idx:      -1,
	}
	return args, nil
}

func openReadFile(args *arguments) ([]string, error) {
	// Открываем файл, имя которого было передано в аргументах
	file, err := os.Open(args.input)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// Инициализируем массив,  в котором будут храниться считанные данные
	rows := make([]string, 0)
	scanner := bufio.NewScanner(file)
	// Считываем построчно данные
	// и если задан флаг i, то соответствующе обрабатываем данные
	if args.i {
		args.match = strings.ToLower(args.match)
		for scanner.Scan() {
			line := strings.ToLower(scanner.Text())
			rows = append(rows, line)
		}
		return rows, nil
	}
	for scanner.Scan() {
		line := scanner.Text()
		rows = append(rows, line)

	}
	return rows, nil
}

// Функция поиска позиции полной строки совпадения
func findFullString(args *arguments, rows []string) {
	for i, v := range rows {
		if v == args.match {
			args.idx = i
			break
		}
	}
}

// Функция поиска позиции переданной строки в данных
func findPosString(args *arguments, rows []string) {
	for i, v := range rows {
		if strings.Contains(v, args.match) {
			args.idx = i
			break
		}
	}
}

// Функция подсчёта количества повторений строки
func countString(args arguments, rows []string) int {
	var count int
	for _, v := range rows {
		if strings.Contains(v, args.match) {
			count++
		}
	}
	return count
}

// Функция удаления строки из данных
func removeString(args arguments, rows []string) {
	// Сценарий при передачи флага F
	if args.F {
		for i, v := range rows {
			if strings.Contains(v, args.match) {
				rows = append(rows[:i], rows[i+1:]...)
			}
		}
		// Сценарий при отсутствии флага F
	} else {
		for i, v := range rows {
			if strings.Contains(v, args.match) {
				rows[i] = strings.ReplaceAll(v, args.match, "")
			}
		}
	}

	fmt.Printf("-v:\n\tДанные после удаления строки '%s':\n%s\n", args.oldMatch, strings.Join(rows, "\n"))
}

func findStringAfter(args arguments, rows []string) {
	fmt.Printf("-A:\n\tСтроки после совпадения с '%s':\n", args.oldMatch)
	rows = rows[args.idx+1:]
	for i, v := range rows {
		fmt.Println(v)
		if i == args.A-1 {
			break
		}
	}
}

func findStringBefore(args arguments, rows []string) {
	fmt.Printf("-B:\n\tСтроки перед совпадением с '%s':\n", args.oldMatch)
	rows = rows[:args.idx]
	for i := len(rows) - 1; i > 0; i-- {
		fmt.Println(rows[i])
		args.B--
		if args.B == 0 {
			break
		}
	}
}
