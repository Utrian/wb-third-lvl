package main

import (
	"fmt"
	"sort"
	"strings"
)

// Функция поиска всех множеств анаграмм по массиву строк
// Требования:
// 1. Входные данные для функции: ссылка на массив, каждый
// элемент которого - слово на русском языке в кодировке
// utf8
// 2. Выходные данные: ссылка на мапу множеств анаграмм
// 3. Ключ - первое встретившееся в словаре слово из
// множества. Значение - ссылка на массив, каждый элемент
// которого, слово из множества.
// 4. Массив должен быть отсортирован по возрастанию.
// 5. Множества из одного элемента не должны попасть в
// результат.
// 6. Все слова должны быть приведены к нижнему регистру.
// 7. В результате каждое слово должно встречаться только один
// раз.

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "одуван"}

	am := *anagram(&words)
	fmt.Println(am)
	// map[листок:0xc0000a80f0 одуван:0xc0000a8120 пятак:0xc0000a80c0]

	for key, value := range am {
		fmt.Println(key, *value)
	}
	// пятак  [пятак пятка тяпка]
	// листок [листок слиток столик]
	// одуван [одуван]
}

// Пункт 1 - вход - ссылка на слайс строк
func anagram(words *[]string) *map[string]*[]string {
	anagrams := make(map[string][]string)

	for _, word := range *words {
		// Пункт 5 - слова меньше 2х букв пропускаем
		if len(word) <= 1 {
			continue
		}
		// Пункт 6 - слова в нижнем регистре
		word := strings.ToLower(word)

		// Приведем слова в единый вид
		letters := strings.Split(word, "")
		sort.Strings(letters)
		key := strings.Join(letters, "")

		// Добавляем новые анаграммы и ВСЕ слова
		// им соответствующие
		if _, ok := anagrams[key]; !ok {
			// 3
			anagrams[key] = []string{word}
		} else {
			anagrams[key] = append(anagrams[key], word)
		}
	}

	// Результирующая мапа
	result := make(map[string]*[]string)

	for _, words := range anagrams {
		// Добавляем ключ в результирующую мапу
		key := words[0]
		sl := make([]string, 0, len(words))
		result[key] = &sl

		// Пункт 4 - сортируем по возврастанию
		sort.Strings(words)

		// Пункт 7 - добавляем слова по одному экземпляру
		var last string
		for _, word := range words {
			if word == last {
				continue
			}
			last = word
			*result[key] = append(*result[key], word)
		}
	}

	// Пункт 2 - выход - ссылка на мапу
	return &result
}
