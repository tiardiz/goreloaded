package firstproject

import (
	"strconv"
	"strings"
	"unicode"
)

// ProcessCommands обрабатывает команды в срезе слов
func ProcessCommands(words []string) []string {
	for i := 0; i < len(words); i++ {
		// Обработка (up, <число>)
		if strings.HasPrefix(words[i], "(up,") {
			if i+1 < len(words) && strings.HasSuffix(words[i+1], ")") {
				numStr := strings.TrimSuffix(words[i+1], ")")
				num, err := strconv.Atoi(numStr)
				if err == nil && i-num >= 0 {
					for j := i - num; j < i; j++ {
						words[j] = strings.ToUpper(words[j])
					}
					words = append(words[:i], words[i+2:]...)
					i--
				}
			}
			continue
		}
		if words[i] == "(up)" {
			if i-1 >= 0 {
				words[i-1] = strings.ToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			}
			continue
		}

		// Обработка (cap)
		if words[i] == "(cap)" {
			if i-1 >= 0 {
				words[i-1] = Capitalize(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			}
			continue
		}

		// Обработка (cap, <число>)
		if strings.HasPrefix(words[i], "(cap,") {
			if i+1 < len(words) && strings.HasSuffix(words[i+1], ")") {
				numStr := strings.TrimSuffix(words[i+1], ")")
				num, err := strconv.Atoi(numStr)
				if err == nil && i-num >= 0 {
					for j := i - num; j < i; j++ {
						words[j] = Capitalize(words[j])
					}
					words = append(words[:i], words[i+2:]...)
					i--
				}
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

		// Обработка (cap, <число>)
		if strings.HasPrefix(words[i], "(low,") {
			if i+1 < len(words) && strings.HasSuffix(words[i+1], ")") {
				numStr := strings.TrimSuffix(words[i+1], ")")
				num, err := strconv.Atoi(numStr)
				if err == nil && i-num >= 0 {
					for j := i - num; j < i; j++ {
						words[j] = strings.ToLower(words[j])
					}
					words = append(words[:i], words[i+2:]...)
					i--
				}
			}
			continue
		}
	}

	return words
}

// Capitalize переписывает слово с заглавной буквы
func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	// Переводим первый символ в верхний регистр
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
