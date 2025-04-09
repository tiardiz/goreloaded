package firstproject

import (
	"strings"
)

func LowCommand(words []string) []string {
	var result []string

	for i := 0; i < len(words); i++ {
		word := words[i]

		if word == "(low)" {
			continue
		}

		if i+1 < len(words) && words[i+1] == "(low)" && words[i+1] != "(cap)" && words[i+1] != "(up)" && words[i+1] != "(bin)" && words[i+1] != "(hex)" {
			result = append(result, strings.ToLower(word))
			i++
		} else {
			result = append(result, word)
		}
	}
	return result
}
