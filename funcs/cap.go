package firstproject

import (
	"unicode"
)

func CapCommand(words []string) []string {
	var result []string

	for i := 0; i < len(words); i++ {
		word := words[i]

		if word == "(cap)" {
			continue
		}

		if i+1 < len(words) && words[i+1] == "(cap)" && words[i+1] != "(low)" && words[i+1] != "(up)" && words[i+1] != "(bin)" && words[i+1] != "(hex)" {
			result = append(result, Capitalize(word))
			i++
		} else {
			result = append(result, word)
		}
	}

	return result
}

// Функция для капитализации первого символа
func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}

	return string(runes)
}
