package firstproject

import (
	"strconv"
)

func BinToDecimal(word string) string {
	num, err := strconv.ParseInt(word, 2, 64)
	if err != nil {
		return word
	}
	return strconv.Itoa(int(num))
}

func BinProcess(words []string) []string {
	var result []string

	for i := 0; i < len(words); i++ {
		word := words[i]

		if word == "(bin)" {
			continue
		}

		if i+1 < len(words) && words[i+1] == "(bin)" && words[i+1] != "(cap)" && words[i+1] != "(up)" && words[i+1] != "(low)" && words[i+1] != "(hex)" {
			result = append(result, BinToDecimal(word))
			i++
		} else {
			result = append(result, word)
		}
	}

	return result
}
