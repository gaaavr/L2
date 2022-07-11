package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

// Паттерн Chain Of Responsibility относится к поведенческим паттернам уровня объекта.
// Паттерн Chain Of Responsibility позволяет избежать привязки объекта-отправителя запроса к объекту-получателю запроса,
// при этом давая шанс обработать этот запрос нескольким объектам.
// Получатели связываются в цепочку, и запрос передается по цепочке, пока не будет обработан каким-то объектом.
// По сути это цепочка обработчиков, которые по очереди получают запрос, а затем решают, обрабатывать его или нет.
// Если запрос не обработан, то он передается дальше по цепочке. Если же он обработан, то паттерн сам решает
// передавать его дальше или нет. Если запрос не обработан ни одним обработчиком, то он просто теряется.

// Применяется:
// 1.Есть более одного объекта, способного обработать запрос, причем настоящий обработчик заранее неизвестен
// и должен быть найден автоматически;
// 2.Необходимо отправить запрос одному из нескольких объектов, не указывая явно, какому именно;
// 3.Набор объектов, способных обработать запрос, должен задаваться динамически.

// Плюсы:
// 1.Компактный синтаксис.
// Минусы:
// 1.Обобщенное назначение.

// Handler представляет собой интерфейс.
type Handler interface {
	SendRequest(message int) string
}

// Каждый отправитель хранит единственную ссылку на начало цепочки,
// а каждый получатель имеет единственную ссылку на своего преемника - последующий элемент в цепочке.

// ConcreteHandlerA обработчик, реализует интерфейс Handler.
type ConcreteHandlerA struct {
	next Handler
}

// SendRequest метод, соответствующий Handler
func (h *ConcreteHandlerA) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Im handler 1"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// ConcreteHandlerB обработчик, реализует интерфейс Handler.
type ConcreteHandlerB struct {
	next Handler
}

// SendRequest метод, соответствующий Handler
func (h *ConcreteHandlerB) SendRequest(message int) (result string) {
	if message == 2 {
		result = "Im handler 2"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// ConcreteHandlerC обработчик, реализует интерфейс Handler.
type ConcreteHandlerC struct {
	next Handler
}

// SendRequest метод, соответствующий Handler
func (h *ConcreteHandlerC) SendRequest(message int) (result string) {
	if message == 3 {
		result = "Im handler 3"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}
