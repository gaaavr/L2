package pattern

import "strings"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

// Суть "Фасада" в разбиении сложной системы на более простые и снижении зависимости подсистем друг от друга.
// Фасад объединяет несколько подсистем в один единый интерфейс.
// Плюсы паттерна в  том, что он позволяет скрыть от клиента элементы сложной системы,
// а также упрощает работу между подсистемами.
// Минусом является возможность излишне переусложнить систему.

// NewMan создаёт мужчину, который включает в себя подсистемы дома, дерева и ребёнка с их методами.
func NewMan() *Man {
	return &Man{
		house: &House{},
		tree:  &Tree{},
		child: &Child{},
	}
}

// Man реализует мужчину и паттерн "Фасад".
type Man struct {
	house *House
	tree  *Tree
	child *Child
}

// Todo возвращает, что мужчина должен сделать.
func (m *Man) Todo() string {
	result := []string{
		m.house.Build(),
		m.tree.Grow(),
		m.child.Born(),
	}
	return strings.Join(result, "\n")
}

// House реализует подсистему "Дом"
type House struct {
}

// Build реализует "стройку" дома.
func (h *House) Build() string {
	return "Build house"
}

// Tree реализует подсистему "Дерево"
type Tree struct {
}

// Grow реализует "выращивание" дерева.
func (t *Tree) Grow() string {
	return "Tree grow"
}

// Child реализует подсистему "Ребёнок"
type Child struct {
}

// Born реализует "рождение" ребёнка.
func (c *Child) Born() string {
	return "Child born"
}
