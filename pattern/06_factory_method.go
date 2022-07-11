package pattern

import "log"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// Паттерн Factory Method относится к порождающим паттернам уровня класса и сфокусирован
// только на отношениях между объектами.
// Паттерн Factory Method полезен, когда система должна оставаться легко расширяемой путем добавления
// объектов новых типов. Этот паттерн является основой для всех порождающих паттернов и может легко
// трансформироваться под нужды системы.
// Паттерн Factory Method применяется для создания объектов с определенным интерфейсом,
// реализации которого предоставляются потомками. Другими словами, есть базовый интерфейс фабрики,
// который говорит, что каждая реализующая его фабрика должна реализовать такой-то метод для создания своих продуктов.

// Плюсы:
// 1.Избавляет класс от привязки к конкретным классам продуктов.
// 2.Выделяет код производства продуктов в одно место, упрощая поддержку кода.
// 3.Упрощает добавление новых продуктов в программу.
// 4.Реализует принцип открытости/закрытости.
// Минусы:
// 1.Может привести к созданию больших параллельных иерархий классов,
// так как для каждого класса продукта надо создать свой подкласс создателя.

// Creator представляет собой интерфейс фабрики.
type Creator interface {
	CreateProduct(action string) Product // Factory Method
}

// Product представляет собой интерфейс продукта.
// Все продукты должны иметь единый интерфейс.
type Product interface {
	Use() string // Каждый продукт может быть использован
}

// ConcreteCreator реализует интерфейс Creator
type ConcreteCreator struct{}

// NewCreator функция конструктор для ConcreteCreator
func NewCreator() Creator {
	return &ConcreteCreator{}
}

// CreateProduct метод фабрики
func (p *ConcreteCreator) CreateProduct(action string) Product {
	var product Product

	switch action {
	case "A":
		product = &ConcreteProductA{action}
	case "B":
		product = &ConcreteProductB{action}
	case "C":
		product = &ConcreteProductC{action}
	default:
		log.Fatalln("Unknown Action")
	}

	return product
}

// ConcreteProductA реализует продукт "A".
type ConcreteProductA struct {
	action string
}

// Use возвращает действие продукта
func (p *ConcreteProductA) Use() string {
	return p.action
}

// ConcreteProductB реализует продукт "B".
type ConcreteProductB struct {
	action string
}

// Use возвращает действие продукта
func (p *ConcreteProductB) Use() string {
	return p.action
}

// ConcreteProductC реализует продукт "C".
type ConcreteProductC struct {
	action string
}

// Use возвращает действие продукта
func (p *ConcreteProductC) Use() string {
	return p.action
}
