package firstproject

import (
	"strings"
	"unicode"
)

func ProcessCommands(words []string) []string {
	for i := 0; i < len(words); i++ {

		if words[i] == "(up)" {
			if i-1 >= 0 {
				words[i-1] = strings.ToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			}
			continue
		}

		if words[i] == "(cap)" {
			if i-1 >= 0 {
				words[i-1] = Capitalize(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			}
			continue
		}

		if words[i] == "(low)" {
			if i-1 >= 0 {
				words[i-1] = strings.ToLower(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			}
			continue
		}

	}

	correctedText := CorrectArticles(words)
	return correctedText
}

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
