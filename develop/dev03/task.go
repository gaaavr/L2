package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type arguments struct {
	k             int
	n, r, u       bool
	input, output string
}

func main() {
	args, err := setFlags()
	if err != nil {
		log.Fatalln(err)
	}
	data, err := openReadFile(args.input)
	if err != nil {
		log.Fatalln(err)
	}
	if args.k < 1 {
		args.k = 0
	} else {
		args.k--
	}
	switch {
	case args.n:
		data = sortOnNumbers(data, args)
	}
	file, err := os.Create("text3.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// преобразуем данные в текст
	lines := make([]string, len(data))
	for i, datum := range data {
		str := strings.Join(datum, " ")
		lines[i] = str
	}
	// записываем текст в файл
	_, err = file.WriteString(strings.Join(lines, "\n"))
	if err != nil {
		panic(err)
	}

}

func setFlags() (arguments, error) {
	var (
		n, r, u bool
		k       int
	)
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющееся строки")
	flag.IntVar(&k, "k", 0, "указание колонки для сортировки")
	flag.Parse()
	input := flag.Arg(0)
	output := flag.Arg(1)

	if input == "" || output == "" {
		return arguments{}, fmt.Errorf("не все имена файлов указаны")
	}
	args := arguments{
		k:      k,
		n:      n,
		r:      r,
		u:      u,
		input:  input,
		output: output,
	}
	return args, nil
}

func openReadFile(f string) ([][]string, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	rows := strings.Split(string(data), "\n")
	matrix := make([][]string, 0, len(rows))
	for _, v := range rows {
		words := strings.Split(v, " ")
		matrix = append(matrix, words)
	}
	return matrix, nil
}

func sortOnColumns(data [][]string, args arguments) [][]string {
	sort.Slice(data, func(i, j int) bool {
		var firstElem, secElem string
		if args.k <= len(data[i]) {
			firstElem = data[i][args.k-1]
		} else {
			firstElem = ""
		}
		if args.k <= len(data[j]) {
			secElem = data[j][args.k-1]
		} else {
			secElem = ""
		}
		return firstElem < secElem
	})
	return data
}

func sortOnNumbers(data [][]string, args arguments) [][]string {
	sort.Slice(data, func(i, j int) bool {
		firstElem, _ := strconv.ParseFloat(getElement(data, i, args.k), 64)
		secElem, _ := strconv.ParseFloat(getElement(data, j, args.k), 64)
		if args.r {
			return firstElem > secElem
		}
		return firstElem < secElem
	})
	return data
}

func getElement(data [][]string, i, k int) string {
	if k < len(data[i]) {
		return data[i][k]
	}
	return ""
}
