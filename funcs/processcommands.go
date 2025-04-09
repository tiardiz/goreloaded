// package firstproject

// func ProcessCommands(words []string) []string {
// 	result := UpCommand(words)
// 	result = LowCommand(result)
// 	result = CapCommand(result)
// 	result = BinProcess(result)
// 	result = HexProcess(result)
// 	result = BinProcess(result)
// 	result = LowCommand(result)
// 	result = LowCommand(result)

// 	correctedText := CorrectArticles(result)
// 	return correctedText
// }

package firstproject

import (
	"strings"
)

func ProcessCommands(words []string) ([]string, int) {
	cnt := 0
	for i := 0; i < len(words); i++ {

		if words[i] == "(up)" {
			if i-1 >= 0 {
				words[i-1] = strings.ToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
				cnt++
			}
			continue
		}

		if words[i] == "(cap)" {
			if i-1 >= 0 {
				words[i-1] = Capitalize(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
				cnt++
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
	return correctedText, cnt
}

func ControlProcess() {

	cnt = 5
	for cnt >= 0 {

		str, cnt2 = ProcessCommands([]string{})
		if cnt2 == 0 {
			break
		}
		cnt++
	}
}
