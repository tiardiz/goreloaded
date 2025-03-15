package firstproject

import (
	"fmt"
	"regexp"
	"strconv"
)

func FindMatches(text string) ([]int, []int) {
	// Регулярное выражение для поиска (cap, d), где d - это число
	re := regexp.MustCompile(`\((cap), (\d+)\)`)

	// Найдем все совпадения с их индексами
	matches := re.FindAllStringSubmatchIndex(text, -1)

	// Срезы для хранения индексов и чисел
	var startIndices []int
	var numbers []int

	// Если совпадения найдены
	if len(matches) > 0 {
		// Перебираем все найденные совпадения
		for _, match := range matches {
			// Индекс начала совпадения
			startIndex := match[0]

			// Извлекаем число из совпадения (между "(cap, " и ")")
			numberStr := text[startIndex+6 : match[1]-1] // 6 — длина "(cap, "
			number, err := strconv.Atoi(numberStr)       // Преобразуем строку в число
			if err != nil {
				fmt.Println("Ошибка при преобразовании строки в число:", err)
			} else {
				// Добавляем индексы и числа в соответствующие срезы
				startIndices = append(startIndices, startIndex)
				numbers = append(numbers, number)
			}
		}
	}

	// Возвращаем срезы с индексами и числами
	return startIndices, numbers
}

// func main() {
// 	// Исходный текст
// 	text := "This is (cap, 2) in the sentence, and here is (cap, 20) with more (cap, 30) data."

// 	// Вызов функции и вывод результата
// 	startIndices, numbers := FindMatches(text)

// 	// Выводим все найденные индексы и числа
// 	for i := range startIndices {
// 		fmt.Printf("Найдено число: %d, Индекс начала: %d\n", numbers[i], startIndices[i])
// 	}
// }
