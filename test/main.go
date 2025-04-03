package main

import (
	"fmt"
	"unicode"
)

func splitWordsAndPunctuation(input string) []string {
	var result []string
	var word []rune

	// Идем по каждому символу в строке
	for _, ch := range input {
		// Если символ — это буква или цифра, добавляем его в слово
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			word = append(word, ch)
		} else {
			// Если нашли знак препинания или пробел
			if len(word) > 0 {
				// Если слово есть, добавляем его в результат
				result = append(result, string(word))
				word = nil // очищаем слово
			}
			if unicode.IsPunct(ch) {
				// Добавляем знак препинания
				result = append(result, string(ch))
			}
		}
	}

	// Добавляем последнее слово, если оно есть
	if len(word) > 0 {
		result = append(result, string(word))
	}

	return result
}

func main() {
	// Исходная строка
	input := "Hello, world! How's it going?"

	// Разделяем строку на слова и знаки препинания
	result := splitWordsAndPunctuation(input)

	// Печатаем результат
	fmt.Println(result)
}
