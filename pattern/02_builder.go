package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Паттерн Builder определяет процесс поэтапного построения сложного объекта.
// После того как будет построена последняя его часть, объект можно использовать.
// Плюсы: использование одного кода для создания различных объектов, изоляция реализуещего сложного кода.
// Минусы: усложнение кода из-за введения дополнительных объектов.

// Builder - интерфейс "строителя".
// Включает в себя методы разработки дизайна автомобиля, сборки его кузова и покраски.
type Builder interface {
	MakeDesign(str string)
	MakeBody(str string)
	MakePainting(str string)
}

// Engineer представляет собой структуру, которая будет "командывать" строителем
type Engineer struct {
	builder Builder
}

// Construct сообщает строителю, что ему делать и в каком порядке
func (d *Engineer) Construct() {
	d.builder.MakeDesign("Design done")
	d.builder.MakeBody("Body done")
	d.builder.MakePainting("Body painted")
}

// ConcreteBuilder реализует интерфейс Builder и взаимодействует со сложным объектом Auto
type ConcreteBuilder struct {
	auto *Auto
}

// MakeDesign делает дизайн кузова
func (b *ConcreteBuilder) MakeDesign(str string) {
	b.auto.Content += str
}

// MakeBody строит кузов автомобиля
func (b *ConcreteBuilder) MakeBody(str string) {
	b.auto.Content += str
}

// MakePainting красит кузов автомобиля
func (b *ConcreteBuilder) MakePainting(str string) {
	b.auto.Content += str
}

// Auto структура, описывающая сложный объект автомобиль
type Auto struct {
	Content string
}

// Show показывает готовый кузов автомобиля
func (a *Auto) Show() string {
	return a.Content
}
