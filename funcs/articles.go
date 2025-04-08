package firstproject

import (
	"strings"
	"unicode"
)

func IsArticle(word string) bool {
	lowerWord := strings.ToLower(word)
	return lowerWord == "a" || lowerWord == "an"
}

func IsConjunction(word string) bool {
	lowerWord := strings.ToLower(word)
	return lowerWord == "or" || lowerWord == "and"
}

func StartsWithVowel(word string) bool {
	if len(word) == 0 {
		return false
	}
	firstChar := unicode.ToLower(rune(word[0]))
	return firstChar == 'a' || firstChar == 'e' || firstChar == 'i' || firstChar == 'o' || firstChar == 'u' || firstChar == 'y'
}

func StartsWithUpperCase(word string) bool {
	if len(word) > 0 && unicode.IsUpper(rune(word[0])) {
		return true
	}
	return false
}

func CorrectArticles(words []string) []string {

	for i := 0; i < len(words)-1; i++ {
		currentWord := words[i]
		nextWord := words[i+1]

		if IsArticle(currentWord) {
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
				// Если артикль был в верхнем регистре (например, "A")
				correctedArticle = Capitalize(correctedArticle) // Используем Capitalize для заглавной буквы
			} else if unicode.IsUpper(rune(currentWord[0])) {
				// Если артикль был с заглавной буквы (например, "An")
				correctedArticle = Capitalize(correctedArticle) // Используем Capitalize для заглавной буквы
			}

			// Заменяем артикль на корректный
			words[i] = correctedArticle
		}
	}
	return words
}
