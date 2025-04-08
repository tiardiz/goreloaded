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
func BinToDecimal(word string) string {
	num, err := strconv.ParseInt(word, 2, 64)
	if err != nil {
		return word
	}
	return strconv.Itoa(int(num))
}

func HexBinProcess(words []string) []string {
	var result []string

	for i := 0; i < len(words); i++ {
		word := words[i]

		if word == "(hex)" || word == "(bin)" {
			continue
		}

		// Если перед словом (hex) или (bin), проверяем, нужно ли преобразовать это слово
		if i+1 < len(words) && words[i+1] == "(hex)" {
			result = append(result, HexToDecimal(word))
			i++
		} else if i+1 < len(words) && words[i+1] == "(bin)" {
			result = append(result, BinToDecimal(word))
			i++
		} else {
			result = append(result, word)
		}
	}

	return result
}
