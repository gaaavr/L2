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
	fmt.Println(unpackingString("f32"))
}

func unpackingString(s string) string {
	var lastSymbol, symbol, unpack string
	for i, v := range s {
		symbol = string(v)
		if i == 0 {
			if _, err := strconv.Atoi(symbol); err == nil {
				return `"" (некорректная строка)`
			}
		}
		digit, err := strconv.Atoi(symbol)
		if err != nil {
			unpack += symbol
			lastSymbol = symbol
		} else {
			for j := 0; j < digit-1; j++ {
				unpack += lastSymbol
			}
		}

	}
	return unpack
}
