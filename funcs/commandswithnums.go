package firstproject

import (
	"math/big"
	"regexp"
	"strings"
)

func CommandsWithNums(words []string) []string {
	re := regexp.MustCompile(`\((cap|up|low),\s*(-?\d+)\)`)

	for i := 0; i < len(words); i++ {
		word := words[i]

		if re.MatchString(word) {
			match := re.FindStringSubmatch(word)
			if len(match) > 0 {
				command := match[1]
				numStr := match[2]

				bigNum := new(big.Int)
				_, ok := bigNum.SetString(numStr, 10)
				if ok && bigNum.Sign() > 0 {
					maxRange := big.NewInt(int64(i))
					if bigNum.Cmp(maxRange) == 1 {
						bigNum = maxRange
					}

					num := int(bigNum.Int64())
					start := i - num
					if start < 0 {
						start = 0
					}
					for j := start; j < i; j++ {
						switch command {
						case "cap":
							words[j] = Capitalize(words[j])
						case "up":
							words[j] = strings.ToUpper(words[j])
						case "low":
							words[j] = strings.ToLower(words[j])
						}
					}
				}
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return words
}
