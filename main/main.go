package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	// Исходный текст
	text := "This is (cap, 2) in the sentence, and here is (cap, 20) with more (cap, 30) data."

	// Регулярное выражение для поиска (cap, d), где d - это число
	re := regexp.MustCompile(`\((cap), (\d+)\)`)

	// Найдем все совпадения с их индексами
	matches := re.FindAllStringSubmatchIndex(text, -1)

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
				// Выводим число и его индекс начала
				fmt.Printf("Найдено число: %d, Индекс начала: %d\n", number, startIndex)
			}
		}
	} else {
		fmt.Println("Совпадений не найдено.")
	}
}
