package firstproject

import (
	"math/big"
	"strconv"
)

func HexToDecimal(word string) string {
	n := new(big.Int)
	_, ok := n.SetString(word, 16)
	if !ok {
		return word // Не удалось распарсить как hex
	}
	return n.String() // Возвращает десятичное представление
}

func BinToDecimal(word string) string {
	num, err := strconv.ParseInt(word, 2, 64)
	if err != nil {
		return word
	}
	return strconv.Itoa(int(num))
}

func BinHexProcess(words []string) []string {

	for i := 0; i < len(words); i++ {
		if words[i] == "(bin)" {
			if i-1 >= 0 {
				words[i-1] = BinToDecimal(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}

		if words[i] == "(hex)" {
			if i-1 >= 0 {
				words[i-1] = HexToDecimal(words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}
	}

	return words
}
