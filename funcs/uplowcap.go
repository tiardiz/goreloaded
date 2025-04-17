package firstproject

import (
	"strings"
	"unicode"
)

func UpLowCapCommands(words []string) []string {
	for i := 0; i < len(words); i++ {

		if words[i] == "(up)" {
			if i-1 >= 0 {
				words[i-1] = strings.ToUpper(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}

		if words[i] == "(cap)" {
			if i-1 >= 0 {
				words[i-1] = Capitalize(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}

		if words[i] == "(low)" {
			if i-1 >= 0 {
				words[i-1] = strings.ToLower(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}
	}
	return words
}

func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}

	runes := []rune(word)
	for i, r := range runes {
		if unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			// Остальные буквы — строчные
			for j := i + 1; j < len(runes); j++ {
				runes[j] = unicode.ToLower(runes[j])
			}
			break
		}
	}
	return string(runes)
}
