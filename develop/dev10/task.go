package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*

Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Требования:
Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера,
программа должна также завершаться. При подключении к несуществующему сервер,
программа должна завершаться через timeout
*/

type join struct {
	host, port string
	timeout    time.Duration
}

func main() {
	j, err := setFlags()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = j.client()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Функция для  распарсивания аргументов командной строки
func setFlags() (join, error) {
	// Объявляем переменные для хранения значений флага
	var (
		h, p string
		t    time.Duration
	)
	// Определяем флаги командной строки и парсим их в переменные
	flag.DurationVar(&t, "timeout", 10*time.Second, "Таймаут на подключение к серверу")
	flag.Parse()

	h = flag.Arg(0)
	p = flag.Arg(1)

	// Инициализируем экземпляр структуры с аргументами командной строки и возвращаем их
	args := join{
		host:    h,
		port:    p,
		timeout: t,
	}
	return args, nil
}

// Метод для подключения к серверу
func (j *join) client() error {
	con, err := net.DialTimeout("tcp", "localhost"+":"+"8080", j.timeout)
	if err != nil {
		return err
	}
	defer func() {
		//закрываем соединение с хостом при выходе из программы
		fmt.Println("Close connection")
		con.Close()
	}()
	// Создаём канал для отслеживания сигналов от ОС, также канал для отлавливания ошибки
	sign, errCh := make(chan os.Signal), make(chan error)
	// При поступлении сигнала от ОС он поступит в канал, в этом случае соединение закроется
	// и программа завершит работу
	signal.Notify(sign, syscall.SIGINT)
	// Читаем сообщение от хоста
	reader, err := bufio.NewReader(con).ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Print(reader)
	scanner := bufio.NewScanner(os.Stdin)
	// Со стандартоного ввода получаем строки и отправляем хосту, который возвращает
	// нам эти строки в верхнем регистре
	go func() {
		for scanner.Scan() {
			_, err := con.Write([]byte(scanner.Text() + "\n"))
			if err != nil {
				errCh <- err
			}
			line, err := bufio.NewReader(con).ReadString('\n')
			if err != nil {
				errCh <- err
			}
			os.Stdout.Write([]byte(line))
		}
	}()

	select {
	case <-sign:
		fmt.Println("Signal received")
		return nil
	case <-errCh:
		return err
	}

	return nil
}
