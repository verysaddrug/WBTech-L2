package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func SetsOfAnagrams(words []string) *map[string][]string {
	anagrams := make(map[string][]string)

	// Приводим все слова к нижнему регистру и сортируем буквы
	for _, word := range words {
		loweredWord := strings.ToLower(word)
		sortedWord := sortStringByCharacter(loweredWord)
		anagrams[sortedWord] = append(anagrams[sortedWord], loweredWord)
	}

	// Формируем результат
	result := make(map[string][]string)
	for _, group := range anagrams {
		if len(group) > 1 {
			sort.Strings(group)
			key := group[0]
			result[key] = group
		}
	}

	return &result
}

// Функция для сортировки букв в строке
func sortStringByCharacter(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func main() {
	words := []string{"пЯтак", "пяТка", "Тяпка", "лИсток", "сЛиток", "стОлик", "беРеза"}
	result := SetsOfAnagrams(words)
	for key, value := range *result {
		fmt.Printf("%s: %v\n", key, value)
	}
}
