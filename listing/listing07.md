Что выведет программа? Объяснить вывод программы.

```go
package main
import (
	"fmt"
	"math/rand"
	"time"
)
func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}
func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Сначала выведутся числа от 1 до 8 в различном порядке из-за рандомайзера по времени
сна пере каждой последующей отдачей числа в канал. Далее буду выводиться 0 в бесконечном
цикле. Несмотря на то, что пишушие данные в канал горутины закрыли каналы после записи,
значения всё равно буду читаться из закрытого канала, но это будут значения по умолчанию.
Для канала с данными типа int значение по умолчанию -0. Для того, чтобы предотвратить
считывание значений по умолчанию из закрытого канала, необходимо в каждом кейсе при 
чтении проверять канал на закрытость с помощью конструкции case v, ok := <-chan, 
где ok =true если канал открыт и false, если закрыт. 
```