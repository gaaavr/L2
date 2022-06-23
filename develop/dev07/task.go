package main

import (
	"fmt"
	"time"
)

/* Реализовать функцию, которая будет объединять один или более done-каналов в single-канал,
если один из его составляющих каналов закроется.
Очевидным вариантом решения могло бы стать выражение при использованием select,
которое бы реализовывало эту связь, однако иногда неизвестно общее число done-каналов,
с которыми вы работаете в рантайме. В этом случае удобнее использовать вызов единственной функции,
которая, приняв на вход один или более or-каналов, реализовывала бы весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}
Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))


*/

func main() {
	// Функция возвращает каналы, которые закрываются при истечении какого-то времени
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	// Засекаем время начала работы
	start := time.Now()
	// В функцию отправляем несколько функций, внутри которых открытые до опредлённого времени каналы
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(99999*time.Millisecond),
		sig(3*time.Second),
		sig(6*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	// Как только один из каналов закрылся, закрылся и общий канал, мэйн горутина завершает работу
	fmt.Printf("fone after %v\n", time.Since(start))
}

// Данная функция принимает на вход несколько каналов
func or(channels ...<-chan interface{}) <-chan interface{} {
	// Создаём два канала. Канал done предназначен для сигнализации
	// текущей функции о завершении анонимных горутин и передачи данных
	// из функции.
	out := make(chan interface{})
	done := make(chan struct{})
	// В цикле в отдельных горутинах ждём данные из каналов, которые переданы функции.
	for i := range channels {
		go func(ch <-chan interface{}) {
			// Когда один из переданных каналов закрывается, то данные из него
			// передаются в канал out, одновременно с этим закрывается
			// канал done. При закрытом done завершаем анонимные горутины.
			select {
			case tmp := <-ch:
				close(done)
				out <- tmp
			case <-done:
				return
			}
		}(channels[i])
	}
	// Разблокируем горутину и отправляем данные из первого закрытого канала
	<-done

	return out
}
