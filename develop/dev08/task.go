package main

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*


Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена
команда выхода (например \quit).
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Получаем полный путь к фалу
	path, err := filepath.Abs("")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(path + ">")
	// При каждом новом вводе определяем команду
	for scanner.Scan() {
		input := scanner.Text()
		args := strings.Split(input, " ")
		switch args[0] {
		// Если передан "cd",то меняем директорию
		case "cd":
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println("Incorrect path")
			}
		// Если передан "pwd",то выводим директорию, в которой находимся
		case "pwd":
			fmt.Println(path)
		// Если передан "echo",то выводим директорию, в которой находимся
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		// Если передан "kill",то завершаем процесс по его pid
		case "kill":
			pid, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println(err)
				break
			}
			prc, err := os.FindProcess(pid)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			err = prc.Kill()
			if err != nil {
				fmt.Println(err.Error())
				break
			}
		// Если передан "ps",то выводим общую информацию по запущенным процессам в виде "имя pid"
		case "ps":
			// Получаем слайс всех запущенных процессов в системе
			prc, err := ps.Processes()
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			var allProcs string
			// В цикле достаём из каждого процесса имя и pid, далее выводим в удобочитаемом виде
			for _, v := range prc {
				allProcs += fmt.Sprintf("%s\t%d\n", v.Executable(), v.Pid())
			}
			fmt.Println(allProcs)
		// При вводе команды "quit" завершаем работу программы
		case "quit":
			fmt.Println("exit")
			os.Exit(0)
		// По дефолту запускаем переданное в аргументе приложение
		default:
			cmd := exec.Command(args[0], args[1:]...)
			//Выводим ошибку, если она есть
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			// Запускаем приложение, указанное в команде
			err := cmd.Run()
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		// Обновляем директорию, в которой находимся
		path, err = filepath.Abs("")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(path + ">")

	}

}
