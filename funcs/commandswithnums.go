package firstproject

import (
	"regexp"
	"strconv"
	"strings"
)

func CommandsWithNums(words []string) []string {
	// Регулярное выражение для поиска команд (cap, <число>), (up, <число>), (low, <число>)
	re := regexp.MustCompile(`\((cap|up|low),\s*(\d+)\)`)

	for i := 0; i < len(words); i++ {
		word := words[i]

		if re.MatchString(word) {
			match := re.FindStringSubmatch(word)
			if len(match) > 0 {
				command := match[1] // Команда (cap, up, low)
				numStr := match[2]  // Число, указывающее количество предыдущих слов
				num, err := strconv.Atoi(numStr)
				if err == nil && i-num >= 0 {

					for j := i - num; j < i; j++ {
						switch command {
						case "cap":
							words[j] = Capitalize(words[j])
						case "up":
							words[j] = strings.ToUpper(words[j])
						case "low":
							words[j] = strings.ToLower(words[j])
						}
					}
					words = append(words[:i], words[i+1:]...)
					i--
				}
			}
		}
	}
	return words
}