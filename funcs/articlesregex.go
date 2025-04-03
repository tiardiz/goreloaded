package firstproject

import (
	"strings"
	"unicode"
)

// Функция для проверки, является ли слово артиклем
func IsArticle(word string) bool {
	lowerWord := strings.ToLower(word)
	return lowerWord == "a" || lowerWord == "an"
}

// Функция для проверки, является ли слово союзом "or" или "and"
func IsConjunction(word string) bool {
	lowerWord := strings.ToLower(word)
	return lowerWord == "or" || lowerWord == "and"
}

// Функция для проверки, начинается ли слово с гласной буквы
func StartsWithVowel(word string) bool {
	if len(word) == 0 {
		return false
	}
	firstChar := unicode.ToLower(rune(word[0]))
	return firstChar == 'a' || firstChar == 'e' || firstChar == 'i' || firstChar == 'o' || firstChar == 'u'
}

// Функция для корректировки артиклей в тексте
func CorrectArticles(text string) string {
	words := strings.Fields(text) // Разбиваем текст на слова

	for i := 0; i < len(words)-1; i++ {
		currentWord := words[i]
		nextWord := words[i+1]

		// Проверяем, является ли текущее слово артиклем
		if IsArticle(currentWord) {
			// Если следующее слово — союз или артикль, то это не артикль
			if IsConjunction(nextWord) || IsArticle(nextWord) {
				continue
			}

			// Определяем, какой артикль должен стоять перед следующим словом
			var correctedArticle string
			if StartsWithVowel(nextWord) {
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
				correctedArticle = Capitalize(correctedArticle) // Используем Capitalize для заглавной буквы
			}

			// Заменяем артикль на корректный
			words[i] = correctedArticle
		}
	}
	return strings.Join(words, " ") // Собираем слова обратно в текст
}
