package main

import (
	"fmt"
	"sort"
	"strings"
)

/* Написать функцию поиска всех множеств анаграмм по словарю.


Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Требования:
Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
Выходные данные: ссылка на мапу множеств анаграмм
Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

*/

func main() {
	// Дан исходный массив слов
	arr := []string{"iiiiiiikkkkkoo", "ikkoooiiiiiiii", "пЯтка", "лиСток", "лиСТок", "слиток", "ПЯТАК", "лишний", "лишний", "слиток", "пятка", "тяпка", "столик"}
	// Печатаем получившийся результат с анаграммами
	fmt.Println(findAnagrams(arr))
}

// Функция принимает на вход слайс строк
func findAnagrams(arr []string) map[string][]string {

	// Иницилизируем три мапы: одну для вывода конечного результата ключом которой является первое слово
	// исходного массива, а значением массив слов-анаграмм ключа со множествами анаграмм.
	// Вторую для проверки повторения слов, третью для проверки существования ключа в первой мапе
	coincidence := make(map[string][]string)
	repeatCheck := make(map[string]struct{})
	keyExists := make(map[string]string)
	// В цикле каждое слово приводим к нижнему регистру, далее проверяем,
	// встречалось ли нам уже такое слово, если да, то сразу переходим к следующей итерации, если нет
	// то добавляем во вторую мапу слово и сортируем символы в слове по возрастанию. Проверяем, существует ли
	// такой ключ в третьей мапе. Если существует, то берём значение по ключу, которое является ключом в основной мапе,
	// и добавляем в массив по этому ключу слово. Если не существует, до добавляем в мапу для проверки ключа.
	for _, v := range arr {
		v = strings.ToLower(v)
		if _, ok := repeatCheck[v]; ok {
			continue
		}
		repeatCheck[v] = struct{}{}
		word := sortWord(v)
		if value, ok := keyExists[word]; ok {
			coincidence[value] = append(coincidence[value], v)
			continue
		}
		keyExists[word] = v
	}
	// В цикле итерируемся по первой мапе.
	// Если длина массива анаграмм меньше 1, то удаляем элемент по этому ключу.
	// Если нет то сортируем массив.
	for i, v := range coincidence {
		if len(coincidence) < 1 {
			delete(coincidence, i)
			continue
		}
		sort.Strings(v)
	}
	// Возвращаем полученную мапу
	return coincidence

}

// Вспомогательная функция для сортировки символов в слове по возрастанию
func sortWord(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
