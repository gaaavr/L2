package main

import (
	"fmt"
	"strconv"
)

/* Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""

Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)


В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.
*/

func main() {
	// Подаём функции на вход строку для распаковки
	unpack, err := unpackingString("qw1e2\\4\\53\\\\6")
	// Если возвращается не нулевая ошибка, выводим её с пустой строкой и выходим из программы
	if err != nil {
		fmt.Println(unpack, err)
		return
	}
	// В случае успешной распаковки выводим полученную строку
	fmt.Println(unpack)
}

// Функция принимает на вход строку для распаковки и возвращает строку и ошибку
func unpackingString(s string) (string, error) {
	// Объявляем строковые переменные для запоминания пердыдущего символа, текущего символа,
	// распакованной строки, представления числа в виде строки, а также
	// булевые переменные для индикатора слэша и escape-последовательности
	var (
		prevSymbol, symbol, unpack, digitStr string
		slash, escape                        bool
	)
	// в цикле проходим по исходной строке и обрабатываем каждый символ строки
	for i, v := range s {
		symbol = string(v)
		// если 1 символ начинается с числа, то возвращаем ошибку
		if i == 0 {
			if _, err := strconv.Atoi(symbol); err == nil {
				return "", fmt.Errorf("(некорректная строка)")
			}
			// иначе начинаем распаковку
			unpack += symbol
			prevSymbol = symbol
			continue
		}
		// если встречается слэш, присваиваем соответствующему индикатору true, а также
		// обозначаем через переменную escape начало escape-последовательности
		if symbol == "\\" {
			// проверка предыдущего значения на слэш
			if prevSymbol == "\\" {
				slash = false
				unpack += prevSymbol
				continue
			}
			slash = true
			escape = true
			prevSymbol = symbol
			continue
		}
		// далее идёт проверка символа на число
		digit, err := strconv.Atoi(symbol)
		if err == nil {
			// если символ является числом и находится в escape-последовательности,
			// то идёт проверка на повтор слэша
			if escape {
				if slash {
					// если предыдущий символ был не слэш, то затираем его новым значением
					// и распаковываем, выключаем индикатор слэша
					if prevSymbol != "\\\\" {
						prevSymbol = symbol
					}
					unpack += prevSymbol
					prevSymbol = symbol
					slash = false
					// в случае выключенного индикатора слэша проверяем предпредыдущий символ на слэш,
					// в положительном случае убавляем количество итераций для распаковки символа на 1
				} else {
					if string(s[i-2]) == "\\" {
						digit -= 1
					}
					// распаковываем символ
					for k := 0; k < digit; k++ {
						unpack += prevSymbol
					}
					// запоминаем текущий символ и переходим к следующему
					prevSymbol = symbol
					continue
				}
				// проверка на последний элемент строки
				if i == len(s)-1 {
					return unpack, nil
				}
				continue
			}
			// если в строке нет escape-последовательности,
			// то запоминаем количество повторений символа через строку в случае,
			// если число имеет больше одного разряда
			digitStr += strconv.Itoa(digit)
			// если символ последний, то переводим число в int и распаковываем последний символ
			if i == len(s)-1 {
				digit, _ = strconv.Atoi(digitStr)
				for j := 0; j < digit-1; j++ {
					unpack += prevSymbol
				}
			}
			// проверка на начало escape-последовательности
			if i < len(s)-1 && string(s[i+1]) == "\\" {
				//если дальше начинается escape-последовательность, то распаковываем последний
				// строковый символ перед слэшем
				digit, _ = strconv.Atoi(digitStr)
				for j := 0; j < digit-1; j++ {
					unpack += prevSymbol
				}
			}
		}
		// если символ не является числом и его количество = 1, то сразу его распаковываем
		if err != nil {
			if digitStr == "" {
				unpack += symbol
				prevSymbol = symbol
				continue

			}
			// если его количество > 1, то в распечатываем в цикле
			digit, _ = strconv.Atoi(digitStr)
			for j := 0; j < digit-1; j++ {
				unpack += prevSymbol
			}
			// распаковываем текущий символ, запоминаем его и обнуляем счётчик
			unpack += symbol
			prevSymbol = symbol
			digitStr = ""

		}

	}
	// возвращаем распакованную строку
	return unpack, nil
}
