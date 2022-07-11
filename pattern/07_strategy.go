package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// Паттерн Strategy относится к поведенческим паттернам уровня объекта.
// Паттерн Strategy определяет набор алгоритмов схожих по роду деятельности, инкапсулирует их в отдельный класс и
// делает их подменяемыми. Паттерн Strategy позволяет подменять алгоритмы без участия клиентов,
// которые используют эти алгоритмы.

// Плюсы:
//	1.Уход от наследования к делегированию
//	2.Изолирует код от алгоритмов
//	3.Горячая замена алгоритмов
//	Минусы:
//	1.Усложняет программу за счет доп классов
//	2.Клиент должен знать в чем различие стратегий

// StrategySort представляет собой интерфейс для сортировочных алгоритмов
type StrategySort interface {
	Sort([]int)
}

// BubbleSort соответствует пузырьковой сортировке
type BubbleSort struct {
}

// Sort сортирует данные пузырьковой сортировкой.
func (s *BubbleSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

// InsertionSort соответствует сортировке вставками
type InsertionSort struct {
}

// Sort сортирует данные через сортировку вставками
func (s *InsertionSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 1; i < size; i++ {
		var j int
		var buff = a[i]
		for j = i - 1; j >= 0; j-- {
			if a[j] < buff {
				break
			}
			a[j+1] = a[j]
		}
		a[j+1] = buff
	}
}

// Context обеспечивает контекст для выполнения стратегии
type Context struct {
	strategy StrategySort
}

// Algorithm заменяет стратегии
func (c *Context) Algorithm(a StrategySort) {
	c.strategy = a
}

// Sort сортирует данные в соответствии с выбранной стратегией
func (c *Context) Sort(s []int) {
	c.strategy.Sort(s)
}
