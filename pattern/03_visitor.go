package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// Паттерн Visitor относится к поведенческим паттернам уровня объекта.
// Паттерн Visitor позволяет обойти набор элементов (объектов) с разнородными интерфейсами,
// а также позволяет добавить новый метод объекта, при этом, не изменяя саму структуру этого объекта.
// Применяется:
// Когда нужно выполнить какую-то операцию над всеми элементами сложной структуры объектов
// Когда новое поведение нужно только для некоторых объектов из существующих
// Плюсы: легко добавлять новые методы для сложных структур.
// Минусы: нарушение инкапсуляции (все типы знают друг о друге), паттерн не оправдан при частой смене иерархии элементов

// Visitor представляет собой интерфейс с несколькими методами, описывающие посетителя
type Visitor interface {
	VisitSushiBar(p *SushiBar) string
	VisitPizzeria(p *Pizzeria) string
	VisitBurgerBar(p *BurgerBar) string
}

// Place представляет собой интерфейс с методом, описывающие место, которое посетит посетитель
type Place interface {
	Accept(v Visitor) string
}

// People реализует интерфейс Visitor
type People struct {
}

// VisitSushiBar метод описывающий посещение заведения
func (v *People) VisitSushiBar(p *SushiBar) string {
	return p.BuySushi()
}

// VisitPizzeria метод описывающий посещение заведения
func (v *People) VisitPizzeria(p *Pizzeria) string {
	return p.BuyPizza()
}

// VisitBurgerBar метод описывающий посещение заведения
func (v *People) VisitBurgerBar(p *BurgerBar) string {
	return p.BuyBurger()
}

// City - структура, содержащая список мест для посещения
type City struct {
	places []Place
}

// Add добавляет место в список мест посещения
func (c *City) Add(p Place) {
	c.places = append(c.places, p)
}

// Accept метод, описывающий посещение всех мест в городе
func (c *City) Accept(v Visitor) string {
	var result string
	for _, p := range c.places {
		result += p.Accept(v)
	}
	return result
}

// SushiBar объект, реализующий интерфейс Place
type SushiBar struct {
}

// Accept - посещение определённого места
func (s *SushiBar) Accept(v Visitor) string {
	return v.VisitSushiBar(s)
}

// BuySushi  покупка продукции
func (s *SushiBar) BuySushi() string {
	return "Buy sushi..."
}

// Pizzeria объект, реализующий интерфейс Place
type Pizzeria struct {
}

// Accept посещение определённого места
func (p *Pizzeria) Accept(v Visitor) string {
	return v.VisitPizzeria(p)
}

// BuyPizza покупка продукции
func (p *Pizzeria) BuyPizza() string {
	return "Buy pizza..."
}

// BurgerBar объект, реализующий интерфейс Place
type BurgerBar struct {
}

// Accept посещение определённого места
func (b *BurgerBar) Accept(v Visitor) string {
	return v.VisitBurgerBar(b)
}

// BuyBurger покупка продукции
func (b *BurgerBar) BuyBurger() string {
	return "Buy burger..."
}
