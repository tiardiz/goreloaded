package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Функция для проверки, является ли слово артиклем
func isArticle(word string) bool {
	lowerWord := strings.ToLower(word)
	return lowerWord == "a" || lowerWord == "an"
}

// Функция для проверки, является ли слово союзом "or" или "and"
func isConjunction(word string) bool {
	lowerWord := strings.ToLower(word)
	return lowerWord == "or" || lowerWord == "and"
}

// Функция для проверки, начинается ли слово с гласной буквы
func startsWithVowel(word string) bool {
	if len(word) == 0 {
		return false
	}
	firstChar := unicode.ToLower(rune(word[0]))
	return firstChar == 'a' || firstChar == 'e' || firstChar == 'i' || firstChar == 'o' || firstChar == 'u'
}

// Функция для корректировки артиклей в тексте
func correctArticles(text string) string {
	words := strings.Fields(text) // Разбиваем текст на слова
	for i := 0; i < len(words)-1; i++ {
		currentWord := words[i]
		nextWord := words[i+1]

		// Проверяем, является ли текущее слово артиклем
		if isArticle(currentWord) {
			// Если следующее слово — союз или артикль, то это не артикль
			if isConjunction(nextWord) || isArticle(nextWord) {
				continue
			}

			// Определяем, какой артикль должен стоять перед следующим словом
			var correctedArticle string
			if startsWithVowel(nextWord) {
				correctedArticle = "an"
			} else {
				correctedArticle = "a"
			}

			// Сохраняем регистр артикля
			if strings.ToUpper(currentWord) == currentWord {
				// Если артикль был в верхнем регистре (например, "AN")
				correctedArticle = strings.ToUpper(correctedArticle)
			} else if unicode.IsUpper(rune(currentWord[0])) {
				// Если артикль был с заглавной буквы (например, "An")
				correctedArticle = strings.Title(correctedArticle)
			}

			// Заменяем артикль на корректный
			words[i] = correctedArticle
		}
	}
	return strings.Join(words, " ") // Собираем слова обратно в текст
}

func main() {
	text := "An cat an pencil AN HOUR a onion An Cat A Apple a a a a A A A A an orange and a banana"
	correctedText := correctArticles(text)
	fmt.Println(correctedText)
}
