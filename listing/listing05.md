Что выведет программа? Объяснить вывод программы.

```go
package main
type customError struct {
	msg string
}
func (e *customError) Error() string {
	return e.msg
}
func test() *customError {
	{
		// do something
	}
	return nil
}
func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Программа выведет error. Сначала err является переменной типа интерфейса error и равна nil.
После выполнения функции test() под капотом переменной в первом поле её структуры содержатся
данные о *customError, так как *customError удовлетворяет интерфейсу error. Следовательно переменная
типа error не будет равна nil.
```