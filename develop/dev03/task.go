package main

import (
	"flag"
	"fmt"
	"log"
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
