package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// Создаём сервер на 80 порту, работающий по протоколу tcp
func main() {
	// Слушаем соединение на 80 порту
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("server is running")
	// В бесконечном цикле при получении соединения сообщаем клиенту о подключении
	for {
		con, err := lis.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		_, err = con.Write([]byte("Welcome to localhost:8080. I change the letter case to capital.\n"))
		if err != nil {
			log.Fatalln(err)
		}
		// Принимаем запросы в виде строк от клиента, преобразуем их в верхний регистр
		// и возвращаем клиенту
		for {
			line, err := bufio.NewReader(con).ReadString('\n')
			// В случае отключения клиента закрываем соединение и ждём соединения со следующим клиентом
			if err != nil {
				con.Close()
				break
			}
			fmt.Println("original string:", line)
			line = "capitalized string: " + strings.ToUpper(line)
			_, err = con.Write([]byte(line))
			if err != nil {
				log.Fatalln(err)
			}
		}

	}
}
