package firstproject

import (
	"strconv"
)

func HexToDecimal(word string) string {
	num, err := strconv.ParseInt(word, 16, 64)
	if err != nil {
		return word
	}
	return strconv.Itoa(int(num))
}

func HexProcess(words []string) []string {
	var result []string

	for i := 0; i < len(words); i++ {
		word := words[i]

		if word == "(hex)" {
			continue
		}

		// Если перед словом (hex) или (bin), проверяем, нужно ли преобразовать это слово
		if i+1 < len(words) && words[i+1] == "(hex)" && words[i+1] != "(cap)" && words[i+1] != "(up)" && words[i+1] != "(bin)" && words[i+1] != "(low)" {
			result = append(result, HexToDecimal(word))
			i++
		} else {
			result = append(result, word)
		}
	}

	return result
}
